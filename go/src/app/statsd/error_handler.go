package statsd

import (
	e "app/events"

	"github.com/quipo/statsd"
)

// NewErrorHandler creates a new statsd CPU message handler
func NewErrorHandler(statsd statsd.StatsdBuffer) GenericHandler {
	return NewGenHandler(
		func(evt e.ProcessEvent) error {
			statsd.Incr(".error", 1)
			return nil
		},
		e.FAILED,
	)
}
