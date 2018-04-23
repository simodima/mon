package main

import (
	"app/conf"
	e "app/events"
	obs "app/observer"
	proc "app/process"
	s "app/statsd"
	"fmt"

	"log"
	"os"
	"sync"

	"github.com/quipo/statsd"
)

var wg sync.WaitGroup

func handlers(sc statsd.StatsdBuffer, c conf.Configuration) []obs.MessageHandler {
	return []obs.MessageHandler{
		obs.DefaultLogger(),
		obs.DefaultMemoryHandler(),
		obs.DefaultCPUHandler(),
		s.NewMemoryHandler(sc),
		s.NewCPUHandler(sc),
		s.NewErrorHandler(sc),
	}
}

func getStatsdClient(c conf.Configuration) statsd.StatsdBuffer {
	return s.NewBufferedClient(fmt.Sprintf("%s:%s", c.StatsdHost, c.StatsdPort), c.ProcessName)
}

func main() {
	args := os.Args[1:]
	conf := conf.ConfigFromEnv(args[0])

	log.Println("Starting main thread")
	statsDClient := getStatsdClient(conf)
	defer statsDClient.Close()

	handlers := handlers(statsDClient, conf)

	processEvents := make(chan e.ProcessEvent)
	startTrigger := make(chan bool)

	wg.Add(1)
	go obs.Start(processEvents, startTrigger, &wg, handlers)

	wg.Add(1)
	p := proc.Init(conf.Tick, processEvents, args)
	go p.Start(startTrigger, &wg)

	wg.Wait()

	log.Println("Ending main thread")
}
