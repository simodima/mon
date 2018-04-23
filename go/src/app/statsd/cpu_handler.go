package statsd

import (
	e "app/events"
	"log"
	"strconv"

	"github.com/quipo/statsd"
)

// NewCPUHandler creates a new statsd CPU message handler
func NewCPUHandler(statsd statsd.StatsdBuffer) GenericHandler {
	return NewGenHandler(
		func(evt e.ProcessEvent) error {
			cpuPercentage, err := strconv.ParseFloat(evt.Data.(string), 64)
			if err != nil {
				log.Fatal(err)
			}

			statsd.Gauge(".cpu", int64(cpuPercentage))

			return nil
		},
		e.CPU,
	)
}
