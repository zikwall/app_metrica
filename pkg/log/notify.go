package log

import (
	"context"
	"sync"

	"github.com/bugsnag/bugsnag-go/v2"
)

var bugsnagEnabled bool

func EnableBugsnagReporter() {
	var once sync.Once
	once.Do(func() {
		bugsnagEnabled = true
	})
}

func ViaNotify(ctx context.Context, err error, meta map[string]interface{}) {
	Warning(err)

	if bugsnagEnabled {
		_ = bugsnag.Notify(err, ctx, bugsnag.MetaData{
			"additional": meta,
		})
	}
}
