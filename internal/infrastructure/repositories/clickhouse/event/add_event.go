package event

import (
	"context"

	"github.com/zikwall/app_metrica/pkg/domain"
)

func (r *Repository) AddEvent(
	ctx context.Context,
	event *domain.EventExtended,
) error {
	r.writer.TryWriteRow(
		domain.RecordFromEvent(event),
	)

	return nil
}
