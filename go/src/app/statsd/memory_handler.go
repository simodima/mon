package statsd

import (
	e "app/events"
	"encoding/json"
	"fmt"

	"github.com/quipo/statsd"
	"github.com/shirou/gopsutil/process"
)

// NewMemoryHandler creates a new statsd Memory message handler
func NewMemoryHandler(statsd statsd.StatsdBuffer) GenericHandler {
	return NewGenHandler(
		func(evt e.ProcessEvent) error {
			var memStats process.MemoryInfoStat
			err := json.Unmarshal([]byte(fmt.Sprint(evt.Data)), &memStats)

			if err != nil {
				return err
			}

			statsd.Gauge(".memory", int64(memStats.RSS))

			return nil
		},
		e.MEMORY,
	)
}
