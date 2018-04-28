package conf

import (
	"app/observer"
	s "app/statsd"

	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/quipo/statsd"
)

// Configuration is the app config
type Configuration struct {
	ProcessName string
	StatsdHost  string
	StatsdPort  string
	Tick        int16
	Statsd      bool
}

func getDefaultConf(procName string) Configuration {
	return Configuration{
		StatsdHost:  "localhost",
		StatsdPort:  "8125",
		ProcessName: procName,
		Tick:        5,
		Statsd:      false,
	}
}

// ConfigFromEnv will create the application config from env variables.
func ConfigFromEnv(procName string) Configuration {
	config := getDefaultConf(procName)
	if err := envconfig.Process("mon", &config); err != nil {
		log.Panic(err)
	}

	return config
}

func getStatsdClient(c Configuration) statsd.StatsdBuffer {
	return s.NewBufferedClient(fmt.Sprintf("%s:%s", c.StatsdHost, c.StatsdPort), c.ProcessName)
}

// GetHandlers will provide the configured process handlers
func GetHandlers(c Configuration) ([]observer.MessageHandler, func()) {
	handlers := []observer.MessageHandler{
		observer.DefaultLogger(),
		observer.DefaultMemoryHandler(),
		observer.DefaultCPUHandler(),
		observer.DefaultNETHandler(),
	}

	close := func() {}

	if c.Statsd == true {
		sc := getStatsdClient(c)
		close = func() {
			defer sc.Close()
		}

		handlers = append(handlers,
			s.NewMemoryHandler(sc),
			s.NewCPUHandler(sc),
			s.NewErrorHandler(sc),
			s.NewNetworkHandler(sc),
		)
	}

	return handlers, close
}
