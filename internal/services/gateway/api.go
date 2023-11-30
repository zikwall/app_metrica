package gateway

import (
	"context"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/mailru/easyjson"
	"github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/pkg/domain"
	"github.com/zikwall/app_metrica/pkg/fiberext"
	"github.com/zikwall/app_metrica/pkg/log"
)

type packet struct {
	data []byte
	ip   string
}

type Handler struct {
	writer *kafka.Writer

	wg    *sync.WaitGroup
	queue chan packet

	poolSize int
}

func (h *Handler) MountRoutes(app *fiber.App) {
	internalV1 := app.Group("/internal/api/v1")
	internalV1.Post("/event", h.Event)
	internalV1.Post("/event-sync", h.EventSync)
}

func (h *Handler) Event(ctx *fiber.Ctx) error {
	body := ctx.Body()

	// we need copy slice, because arrays, yes
	dst := make([]byte, len(body))
	copy(dst, body)

	// non-blocking, asynchronously write to queue and Kafka
	h.queue <- packet{
		data: dst,
		ip:   fiberext.RealIP(ctx),
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (h *Handler) EventSync(ctx *fiber.Ctx) error {
	if err := h.handle(ctx.Context(), packet{
		data: ctx.Body(),
		ip:   fiberext.RealIP(ctx),
	}); err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (h *Handler) Run(ctx context.Context) {
	for i := 1; i <= h.poolSize; i++ {
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

	var err error

	for {
		select {
		case <-ctx.Done():
			return
		case b := <-h.queue:
			if err = h.handle(ctx, b); err != nil {
				log.Warningf("producer proc: handle %s", err)
			}
		}
	}
}

func (h *Handler) handle(ctx context.Context, p packet) error {
	var (
		err   error
		event = &domain.Event{}
	)
	if err = easyjson.Unmarshal(p.data, event); err != nil {
		return err
	}

	// optimized read? - need test
	//
	// var b bytes.Buffer
	// sizedBuffer := bufio.NewWriterSize(&b, len(body)+1)
	// if _, err = easyjson.MarshalToWriter(event, sizedBuffer); err != nil {
	//   return err
	// }

	exEvent := domain.ExtendEvent(event)
	exEvent.IP = p.ip

	var nextForwardingBytes []byte
	if nextForwardingBytes, err = easyjson.Marshal(exEvent); err != nil {
		return err
	}

	return h.writer.WriteMessages(ctx, kafka.Message{
		Value: nextForwardingBytes,
	})
}

func NewHandler(writer *kafka.Writer, poolSize int) *Handler {
	return &Handler{
		writer:   writer,
		wg:       &sync.WaitGroup{},
		queue:    make(chan packet, poolSize+10000),
		poolSize: poolSize,
	}
}
