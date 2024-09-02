package event

import (
	"fmt"
	"strings"
	"time"
)

//easyjson:skip
type EventDatetime struct {
	time.Time
}

const (
	dtLayout       = "2006-01-02 15:04:05"
	dtLayoutWithTz = "2006-01-02 15:04:05 -07:00"

	lenWithTz = 26
)

func (ct *EventDatetime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	if len(s) == lenWithTz {
		ct.Time, err = time.Parse(dtLayoutWithTz, s)
		return
	}
	ct.Time, err = time.Parse(dtLayout, s)
	return
}

func (ct *EventDatetime) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(dtLayoutWithTz))), nil
}

func (ct *EventDatetime) Validate() error {
	minDatetime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	maxDatetime := time.Date(2106, 2, 7, 6, 28, 15, 0, time.UTC)

	if ct.IsZero() {
		return fmt.Errorf("invalid event_datetime: zero value")
	}
	if ct.Before(minDatetime) || ct.After(maxDatetime) {
		return fmt.Errorf("invalid event_datetime: out of range [%s, %s]", minDatetime, maxDatetime)
	}
	return nil
}
