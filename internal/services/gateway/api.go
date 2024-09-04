package gateway

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/zikwall/app_metrica/internal/metrics"
	"github.com/zikwall/app_metrica/internal/services/gateway/eventbus"
	"github.com/zikwall/app_metrica/pkg/domain/mediavitrina"
	"github.com/zikwall/app_metrica/pkg/fiberext"
)

type EventBus interface {
	SendEvent(event eventbus.Event)
	DebugMessage(event eventbus.Event) error
}

type Handler struct {
	ownBus   EventBus
	mediaBus EventBus

	metrics        *metrics.Metrics
	metricsVitrina *metrics.MetricsVitrina
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
// @Param			data	body	event.Event	true	"request debug event"
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

	err := h.ownBus.DebugMessage(
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
// @Param			data	body	event.Event	true	"request event"
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
	h.ownBus.SendEvent(
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
// @Param			data	body	event.Events	true	"request events"
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
	h.ownBus.SendEvent(
		eventbus.NewEvent(dst, utils.CopyString(fiberext.RealIP(ctx)), time.Now(), eventbus.EventTypeBatch),
	)

	h.metrics.IncRequests(true)
	h.metrics.RequestsDuration(start, true)

	return ctx.SendStatus(http.StatusNoContent)
}

// @Summary		Receive eventMedia fields
// @Description	Method receive messages and send to queue
// @Tags			Events
// @Accept			json
// @Param			data	body	mediavitrina.MediaVitrina	true	"request events"
// @Produce		json
// @Success		201	{object}	string	"no content"
// @Router			/internal/api/v1/event/mediavitrina [get]
func (h *Handler) eventMedia(ctx *fiber.Ctx) error {
	mv := mediavitrina.QueryParametersToEntity(ctx.Queries())

	dst, err := mv.Bytes()
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	// non-blocking, asynchronously write to queue and EventBus
	h.mediaBus.SendEvent(
		eventbus.NewEvent(dst, utils.CopyString(fiberext.RealIP(ctx)), time.Now(), eventbus.EventTypeInline),
	)

	return ctx.Status(http.StatusOK).Send(dst)
}

func NewHandler(
	ownBus EventBus,
	mediaBus EventBus,
	metrics *metrics.Metrics,
	metricsVitrina *metrics.MetricsVitrina,
) *Handler {
	return &Handler{
		ownBus:         ownBus,
		mediaBus:       mediaBus,
		metrics:        metrics,
		metricsVitrina: metricsVitrina,
	}
}
