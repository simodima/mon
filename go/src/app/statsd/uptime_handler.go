package statsd

import (
	e "app/events"
	"fmt"
	"time"

	"github.com/quipo/statsd"
)

// NewUptimeHandler creates a new statsd Memory message handler
func NewUptimeHandler(statsd statsd.StatsdBuffer) GenericHandler {
	return NewGenHandler(
		func(evt e.ProcessEvent) error {
			duration, err := time.ParseDuration(fmt.Sprintf("%ss", evt.Data))
			if err != nil {
				return err
			}

			statsd.Gauge(".uptime", int64(duration.Seconds()))

			return nil
		},
		e.UPTIME,
	)
}
