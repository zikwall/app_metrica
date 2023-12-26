package gateway

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mailru/easyjson"
	"github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/config"
	"github.com/zikwall/app_metrica/pkg/domain"
	"github.com/zikwall/app_metrica/pkg/fiberext"
	"github.com/zikwall/app_metrica/pkg/log"
)

type packet struct {
	data []byte
	ip   string
	t    time.Time
}

type Handler struct {
	wg    *sync.WaitGroup
	queue chan packet

	opt *config.Config
}

func (h *Handler) MountRoutes(app *fiber.App) {
	internalV1 := app.Group("/internal/api/v1")
	internalV1.Post("/event", h.Event)
}

func (h *Handler) Event(ctx *fiber.Ctx) error {
	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	// non-blocking, asynchronously write to queue and Kafka
	h.queue <- packet{
		data: dst,
		ip:   fiberext.RealIP(ctx),
		t:    time.Now(),
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (h *Handler) Run(ctx context.Context) {
	for i := 1; i <= h.opt.Internal.HandlerProcSize; i++ {
		h.wg.Add(1)
		go func(number int) {
			defer h.wg.Done()

			h.proc(ctx, number)
		}(i)
	}

	h.wg.Wait()
}

func (h *Handler) proc(ctx context.Context, number int) {
	log.Infof("run producer proc: %d", number)
	defer log.Infof("stop producer proc: %d", number)

	writer := &kafka.Writer{
		Addr:            kafka.TCP(h.opt.KafkaWriter.Brokers...),
		Topic:           h.opt.KafkaWriter.Topic,
		Balancer:        &kafka.LeastBytes{},
		MaxAttempts:     h.opt.KafkaWriter.MaxAttempts,
		WriteBackoffMin: h.opt.KafkaWriter.WriteBackoffMin,
		WriteBackoffMax: h.opt.KafkaWriter.WriteBackoffMax,
		BatchSize:       h.opt.KafkaWriter.BatchSize,
		BatchBytes:      h.opt.KafkaWriter.BatchBytes,
		BatchTimeout:    h.opt.KafkaWriter.BatchTimeout,
		ReadTimeout:     h.opt.KafkaWriter.ReadTimeout,
		WriteTimeout:    h.opt.KafkaWriter.WriteTimeout,
	}

	defer func() {
		if err := writer.Close(); err != nil {
			log.Warningf("failed to drop Kafka producer: %s", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case b := <-h.queue:
			next, err := h.handle(b)
			if err != nil {
				log.Warningf("producer proc: handle %s", err)
				continue
			}

			if err = writer.WriteMessages(ctx, kafka.Message{
				Value: next,
			}); err != nil {
				log.Warningf("producer proc: handle %s", err)
			}
		}
	}
}

func (h *Handler) handle(p packet) ([]byte, error) {
	var (
		err   error
		event = &domain.Event{}
	)
	if err = easyjson.Unmarshal(p.data, event); err != nil {
		return nil, err
	}

	exEvent := domain.ExtendEvent(event, time.Now(), p.t)
	exEvent.IP = p.ip

	var nextForwardingBytes []byte
	if nextForwardingBytes, err = easyjson.Marshal(exEvent); err != nil {
		return nil, err
	}
	return nextForwardingBytes, nil
}

func NewHandler(opt *config.Config) *Handler {
	return &Handler{
		wg:    &sync.WaitGroup{},
		queue: make(chan packet, opt.Internal.HandlerProcSize+10000),
		opt:   opt,
	}
}
