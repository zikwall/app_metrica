package event

import (
	"context"

	"github.com/zikwall/app_metrica/pkg/domain/event"
)

func (r *Repository) AddEvent(
	_ context.Context,
	event *event.EventExtended,
) error {
	r.writer.TryWriteRow(
		recordFromEvent(event),
	)
	return nil
}
