package eventbus

import "time"

type EventType uint8

const (
	EventTypeInline EventType = iota
	EventTypeBatch
)

type Event struct {
	Data []byte
	IP   string
	T    time.Time

	Evt EventType
}

func NewEvent(data []byte, ip string, received time.Time, evt EventType) Event {
	return Event{
		Data: data,
		IP:   ip,
		T:    received,
		Evt:  evt,
	}
}
