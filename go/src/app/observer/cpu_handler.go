package observer

import (
	e "app/events"
	"fmt"
	"log"
)

func DefaultCPUHandler() CPUHandler {
	return CPUHandler{
		func(msg string) {
			log.Printf("CPU usage %s %%", msg)
		},
	}
}

type CPUHandler struct {
	handle func(string)
}

func (h CPUHandler) Supports(evtName string) bool {
	return evtName == e.CPU
}

func (h CPUHandler) Handle(evt e.ProcessEvent) error {
	h.handle(fmt.Sprint(evt.Data))

	return nil
}
