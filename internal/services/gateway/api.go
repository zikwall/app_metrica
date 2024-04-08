package gateway

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/zikwall/app_metrica/internal/eventbus"
	"github.com/zikwall/app_metrica/pkg/fiberext"
)

type EventBus interface {
	SendEvent(event eventbus.Event)
	DebugMessage(event eventbus.Event) error
}

type Handler struct {
	bus EventBus
}

func (h *Handler) MountRoutes(app *fiber.App) {
	internalV1 := app.Group("/internal/api/v1")
	internalV1.Post("/event", h.event)
	internalV1.Post("/event-batch", h.eventBatch)

	internalV1.Post("/event-debug", h.eventDebug)
}

type ErrorMessage struct {
	Error       string `json:"error"`
	RequestBody string `json:"request_body"`
}

func (h *Handler) eventDebug(ctx *fiber.Ctx) error {
	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	err := h.bus.DebugMessage(eventbus.NewEvent(dst, fiberext.RealIP(ctx), time.Now(), eventbus.EventTypeInline))
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(ErrorMessage{
			Error:       err.Error(),
			RequestBody: string(body),
		})
	}
	return ctx.SendStatus(http.StatusAccepted)
}

func (h *Handler) event(ctx *fiber.Ctx) error {
	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	// non-blocking, asynchronously write to queue and EventBus
	h.bus.SendEvent(
		eventbus.NewEvent(dst, fiberext.RealIP(ctx), time.Now(), eventbus.EventTypeInline),
	)
	return ctx.SendStatus(http.StatusNoContent)
}

func (h *Handler) eventBatch(ctx *fiber.Ctx) error {
	// Returned value is only valid within the handler. Do not store any references.
	// Make copies or use the Immutable setting instead.
	body := ctx.Body()

	// See: https://docs.gofiber.io/#zero-allocation
	dst := make([]byte, len(body))
	copy(dst, body)

	// non-blocking, asynchronously write to queue and EventBus
	h.bus.SendEvent(
		eventbus.NewEvent(dst, fiberext.RealIP(ctx), time.Now(), eventbus.EventTypeBatch),
	)
	return ctx.SendStatus(http.StatusNoContent)
}

func NewHandler(bus EventBus) *Handler {
	return &Handler{
		bus: bus,
	}
}
