package xerror

import (
	"errors"
	"fmt"
)

var (
	ErrMaxMindWithGeo = errors.New("max_mind is empty when with_geo is true")
)

type ErrPacket struct {
	err  error
	data string
}

func NewErrPacket(err error, data string) error {
	return &ErrPacket{
		err:  err,
		data: data,
	}
}

func (e *ErrPacket) Error() string {
	return fmt.Sprintf("%s, %s", e.err.Error(), e.data)
}

func (e *ErrPacket) Unwrap() error {
	return e.err
}
