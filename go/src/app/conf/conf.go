package conf

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Configuration is the app config
type Configuration struct {
	ProcessName string
	StatsdHost  string
	StatsdPort  string
	Tick        int16
}

func getDefaultConf(procName string) Configuration {
	return Configuration{
		StatsdHost:  "localhost",
		StatsdPort:  "8125",
		ProcessName: procName,
		Tick:        5,
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
