package xerror

import "errors"

var (
	ErrMaxMindWithGeo = errors.New("max_mind is empty when with_geo is true")
)
