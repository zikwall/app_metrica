package own

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mailru/easyjson"

	"github.com/zikwall/app_metrica/internal/services/gateway/eventbus"
	"github.com/zikwall/app_metrica/pkg/domain/event"
	"github.com/zikwall/app_metrica/pkg/xerror"
)

func HandleBytes(e eventbus.Event) ([]byte, error) {
	var (
		err error
		evt = &event.Event{}
	)
	if err = easyjson.Unmarshal(e.Data, evt); err != nil {
		return nil, xerror.NewErrPacket(
			fmt.Errorf("unmarshal json object: %w", err),
			string(e.Data),
		)
	}

	exEvent := event.ExtendEvent(evt, time.Now(), e.T)
	exEvent.IP = e.IP

	var bytes []byte
	if bytes, err = easyjson.Marshal(exEvent); err != nil {
		return nil, err
	}
	return bytes, nil
}

func HandleMultipleBytes(e eventbus.Event) ([][]byte, error) {
	var (
		err    error
		events event.Events
	)
	if err = easyjson.Unmarshal(e.Data, &events); err != nil {
		return nil, xerror.NewErrPacket(
			fmt.Errorf("unmarshal json object: %w", err),
			string(e.Data),
		)
	}

	bytes := make([][]byte, len(events))
	for i := range events {
		exEvent := event.ExtendEvent(events[i], time.Now(), e.T)
		exEvent.IP = e.IP

		var extBytes []byte
		if extBytes, err = easyjson.Marshal(exEvent); err != nil {
			return nil, err
		}

		bytes[i] = extBytes
	}

	return bytes, nil
}

func DebugMessage(ev eventbus.Event) error {
	var (
		err error
		evt = &event.Event{}
	)
	if err = json.Unmarshal(ev.Data, evt); err != nil {
		return err
	}

	return evt.EventDatetime.Validate()
}
