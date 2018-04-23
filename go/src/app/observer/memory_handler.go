package observer

import (
	e "app/events"
	"fmt"
	"log"
)

func DefaultMemoryHandler() MemoryHandler {
	return MemoryHandler{
		func(msg string) {
			log.Printf("MEMORY %s", msg)
		},
	}
}

type MemoryHandler struct {
	handle func(string)
}

func (l MemoryHandler) Supports(evtName string) bool {
	return evtName == e.MEMORY
}

func (l MemoryHandler) Handle(evt e.ProcessEvent) error {
	l.handle(fmt.Sprint(evt.Data))

	return nil
}
