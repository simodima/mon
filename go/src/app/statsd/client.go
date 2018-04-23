package statsd

import (
	"log"
	"os"
	"time"

	"github.com/quipo/statsd"
)

// NewBufferedClient gives a buffered statsd client
// the buffe will be flushed every 2 secs
func NewBufferedClient(connString string, prefix string) statsd.StatsdBuffer {
	statsdclient := statsd.NewStatsdClient(connString, prefix)
	err := statsdclient.CreateSocket()
	if nil != err {
		log.Println(err)
		os.Exit(1)
	}

	interval := time.Second * 2 // aggregate stats and flush every 2 seconds
	stats := statsd.NewStatsdBuffer(interval, statsdclient)

	return *stats
}
