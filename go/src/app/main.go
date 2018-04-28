package main

import (
	"app/conf"
	"app/events"
	"app/observer"
	"app/process"

	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	args := os.Args[1:]
	config := conf.ConfigFromEnv(args[0])
	handlers, close := conf.GetHandlers(config)

	log.Println("Starting main thread")

	processEvents := make(chan events.ProcessEvent)
	startTrigger := make(chan bool)

	wg.Add(1)
	go observer.Start(processEvents, startTrigger, &wg, handlers)

	wg.Add(1)
	p := process.Init(config.Tick, processEvents, args)
	go p.Start(startTrigger, &wg)

	wg.Wait()
	close()

	log.Println("Ending main thread")
}
