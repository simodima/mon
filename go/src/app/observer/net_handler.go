package observer

import (
	e "app/events"
	"fmt"
	"log"
)

func DefaultNETHandler() NETHandler {
	return NETHandler{
		func(msg string) {
			log.Printf("NET usage %s", msg)
		},
	}
}

type NETHandler struct {
	handle func(string)
}

func (h NETHandler) Supports(evtName string) bool {
	return evtName == e.NETWORK
}

func (h NETHandler) Handle(evt e.ProcessEvent) error {
	h.handle(fmt.Sprint(evt.Data))

	return nil
}
