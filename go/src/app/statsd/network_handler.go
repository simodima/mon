package statsd

import (
	e "app/events"
	"encoding/json"
	"fmt"

	"github.com/quipo/statsd"
	"github.com/shirou/gopsutil/net"
)

// NewNetworkHandler creates a new statsd Memory message handler
func NewNetworkHandler(statsd statsd.StatsdBuffer) GenericHandler {
	return NewGenHandler(
		func(evt e.ProcessEvent) error {
			var netStats net.IOCountersStat
			err := json.Unmarshal([]byte(fmt.Sprint(evt.Data)), &netStats)

			if err != nil {
				return err
			}

			statsd.Gauge(".net.byte_sent", int64(netStats.BytesSent))
			statsd.Gauge(".net.byte_recv", int64(netStats.BytesRecv))

			return nil
		},
		e.NETWORK,
	)
}
