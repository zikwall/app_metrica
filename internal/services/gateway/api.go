package gateway

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/zikwall/app_metrica/internal/eventbus"
	"github.com/zikwall/app_metrica/internal/metrics"
	"github.com/zikwall/app_metrica/pkg/fiberext"
)

type EventBus interface {
	SendEvent(event eventbus.Event)
	DebugMessage(event eventbus.Event) error
}

type Handler struct {
	bus     EventBus
	metrics *metrics.Metrics
}

// MountRoutes godoc
//
//	@title			OwnMetrics aka AppMetrica gateway service
//	@version		1.0
//	@description	Under construct
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	a.kapitonov@limehd.tv
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			lm.limehd.tv
//	@BasePath		/
func (h *Handler) MountRoutes(app *fiber.App) {
	internalV1 := app.Group("/internal/api/v1")
	internalV1.Post("/event", h.event)
	internalV1.Post("/event-batch", h.eventBatch)

	internalV1.Post("/event-debug", h.eventDebug)

	internalV1.Get("/event/mediavitrina", h.eventMedia)
}

type ErrorMessage struct {
	Error       string `json:"error"`
	RequestBody string `json:"request_body"`
}

type DebugMessage struct {
	RealIP        string   `json:"real_ip"`
	IP            string   `json:"ip"`
	IPs           []string `json:"ips"`
	XRealIP       string   `json:"x_real_ip"`
	XForwardedFor string   `json:"x_forwarded_for"`
}

// @Summary		Debug event fields
// @Description	Method check event message is valid or not
// @Tags			Events
// @Accept			json
// @Param			data	body	domain.Event	true	"request debug event"
// @Produce		json
// @Success		202	{object}	DebugMessage	"ok"
// @Failure		422	{object}	ErrorMessage
// @Router			/internal/api/v1/event-debug [post]
func (h *Handler) eventDebug(ctx *fiber.Ctx) error {
	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	err := h.bus.DebugMessage(
		eventbus.NewEvent(dst, utils.CopyString(fiberext.RealIP(ctx)), time.Now(), eventbus.EventTypeInline),
	)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(ErrorMessage{
			Error:       err.Error(),
			RequestBody: string(body),
		})
	}
	return ctx.Status(http.StatusAccepted).JSON(DebugMessage{
		RealIP:        fiberext.RealIP(ctx),
		IP:            ctx.IP(),
		IPs:           ctx.IPs(),
		XRealIP:       utils.CopyString(ctx.Get("X-Real-IP")),
		XForwardedFor: utils.CopyString(ctx.Get("X-Forwarded-For")),
	})
}

// @Summary		Receive event fields
// @Description	Method receive message and send to queue
// @Tags			Events
// @Accept			json
// @Param			data	body	domain.Event	true	"request event"
// @Produce		json
// @Success		201	{object}	string	"no content"
// @Router			/internal/api/v1/event [post]
func (h *Handler) event(ctx *fiber.Ctx) error {
	start := time.Now()

	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	// non-blocking, asynchronously write to queue and EventBus
	h.bus.SendEvent(
		eventbus.NewEvent(dst, utils.CopyString(fiberext.RealIP(ctx)), time.Now(), eventbus.EventTypeInline),
	)

	h.metrics.IncRequests(false)
	h.metrics.RequestsDuration(start, false)

	return ctx.SendStatus(http.StatusNoContent)
}

// @Summary		Receive event fields with batch mode
// @Description	Method receive messages and send to queue
// @Tags			Events
// @Accept			json
// @Param			data	body	domain.Events	true	"request events"
// @Produce		json
// @Success		201	{object}	string	"no content"
// @Router			/internal/api/v1/event-batch [post]
func (h *Handler) eventBatch(ctx *fiber.Ctx) error {
	start := time.Now()

	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	// non-blocking, asynchronously write to queue and EventBus
	h.bus.SendEvent(
		eventbus.NewEvent(dst, utils.CopyString(fiberext.RealIP(ctx)), time.Now(), eventbus.EventTypeBatch),
	)

	h.metrics.IncRequests(true)
	h.metrics.RequestsDuration(start, true)

	return ctx.SendStatus(http.StatusNoContent)
}

func (h *Handler) eventMedia(ctx *fiber.Ctx) error {
	queries := ctx.Queries()

	// temporary
	file, err := os.OpenFile("/tmp/media_queries", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString("\n")
	_, err = file.WriteString(time.Now().String())

	b, err := json.Marshal(queries)
	if err != nil {
		return err
	}

	_, err = file.Write(b)

	return ctx.Status(http.StatusOK).JSON(queries)
}

func NewHandler(bus EventBus, metrics *metrics.Metrics) *Handler {
	return &Handler{
		bus:     bus,
		metrics: metrics,
	}
}
