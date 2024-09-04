package mediavitrina

import (
	"context"
	"fmt"

	"github.com/mailru/easyjson"
	kafkaLib "github.com/segmentio/kafka-go"

	"github.com/zikwall/app_metrica/pkg/domain/mediavitrina"
)

type EventRepository interface {
	AddMediaEvent(
		ctx context.Context,
		event *mediavitrina.MediaVitrina,
	) error
}

type Metrics interface {
	IncConsumerError(err error)
}

type Handler struct {
	eventRepository EventRepository
	metrics         Metrics
}

func (h *Handler) Handle(ctx context.Context, message kafkaLib.Message) error {
	evt := &mediavitrina.MediaVitrina{}
	if err := easyjson.Unmarshal(message.Value, evt); err != nil {
		h.metrics.IncConsumerError(err)

		return fmt.Errorf("consumer proc: unmarshal json %w", err)
	}

	if err := h.eventRepository.AddMediaEvent(ctx, evt); err != nil {
		return fmt.Errorf("consumer proc: add event: %w", err)
	}

	return nil
}

func New(eventRepository EventRepository, metrics Metrics) *Handler {
	return &Handler{
		eventRepository: eventRepository,
		metrics:         metrics,
	}
}
