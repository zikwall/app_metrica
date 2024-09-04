package mediavitrina

import (
	"github.com/mailru/easyjson"

	"github.com/zikwall/app_metrica/internal/services/gateway/eventbus"
	"github.com/zikwall/app_metrica/pkg/domain/mediavitrina"
)

func HandleBytes(ev eventbus.Event) ([]byte, error) {
	return ev.Data, nil
}

func HandleMultipleBytes(_ eventbus.Event) ([][]byte, error) {
	panic("unsupported")
}

func DebugMessage(ev eventbus.Event) error {
	var (
		err error
		evt = &mediavitrina.MediaVitrina{}
	)
	if err = easyjson.Unmarshal(ev.Data, evt); err != nil {
		return err
	}

	return evt.EventDatetime.Validate()
}
