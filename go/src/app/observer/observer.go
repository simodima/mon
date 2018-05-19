package observer

import (
	e "app/events"
	"log"
	"sync"
	"time"
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
				errResult := make(chan error, 1)
				go func() { errResult <- h.Handle(event) }()
				select {
				case err := <-errResult:
					if err != nil {
						logError(err)
					}
				case <-time.After(time.Second * time.Duration(1)):
					log.Printf("Event processing timed out for %s", event.Name)
				}
			}
		}

		if event.Name == e.DONE || event.Name == e.FAILED {
			log.Printf("Received %s from subproces. Eding observer", event.Name)
			return
		}
	}
}
