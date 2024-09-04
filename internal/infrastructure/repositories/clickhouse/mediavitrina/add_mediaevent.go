package mediavitrina

import (
	"context"

	"github.com/zikwall/app_metrica/pkg/domain/mediavitrina"
)

func (r *Repository) AddMediaEvent(
	_ context.Context,
	event *mediavitrina.MediaVitrina,
) error {
	r.writer.TryWriteRow(
		recordFromEvent(event),
	)
	return nil
}
