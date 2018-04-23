package observer

import (
	e "app/events"
	"log"
	"sync"
)

func logError(err error) {
	log.Print(err)
}

// Start a new observer
func Start(pEvents chan e.ProcessEvent, startTrigger chan bool, wg *sync.WaitGroup, handlers []MessageHandler) {
	defer wg.Done()

	log.Println("Sending start subprocess signal")
	startTrigger <- true
	log.Println("Waiting event from subprocess")
	for {
		event := <-pEvents
		for _, h := range handlers {
			if h.Supports(event.Name) {
				if err := h.Handle(event); err != nil {
					logError(err)
				}

			}
		}

		if event.Name == e.DONE || event.Name == e.FAILED {
			log.Printf("Received %s from subproces. Eding observer", event.Name)
			return
		}
	}
}
