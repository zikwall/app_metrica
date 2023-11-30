package xerror

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	status     int
	text       string
	additional string
}

func (e *HTTPError) Error() string {
	if e.additional == "" {
		return fmt.Sprintf("http error: [%d] %s", e.status, e.text)
	}
	return fmt.Sprintf("http error: [%d] %s (%s)", e.status, e.text, e.additional)
}

func (e *HTTPError) Code() int {
	return e.status
}

func (e *HTTPError) Message() string {
	return e.text
}

func newHTTPError(status int, message string, err ...error) *HTTPError {
	e := &HTTPError{
		status: status,
		text:   message,
	}
	if len(err) > 0 {
		e.additional = err[0].Error()
	}
	return e
}

// NewBadRequest dynamic bad request error with additional error information
func NewBadRequest(err error) *HTTPError {
	return newHTTPError(http.StatusBadRequest, "400 Bad Request", err)
}

// static errors
var (
	ErrBadRequest = &HTTPError{
		status: http.StatusBadRequest,
		text:   "400 Bad Request",
	}
	ErrNotFoundEpgIsZero = &HTTPError{
		status: http.StatusNotFound,
		text:   "404 Not Found: epg was not found by ID or not available for your application and region",
	}
)
