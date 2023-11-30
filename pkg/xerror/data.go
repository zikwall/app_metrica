package xerror

import "errors"

var (
	ErrRowNotFound        = errors.New("row not found")
	ErrEpgIDIsNull        = errors.New("epg_id is null")
	ErrMaxIdleConns       = errors.New("max_idle_conns must be greater than zero")
	ErrMaxOpenConns       = errors.New("max_open_conns must be greater than zero")
	ErrMaxConnMaxLifetime = errors.New("max_conn_max_lifetime must be greater than zero")
)
