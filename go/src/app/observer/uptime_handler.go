package observer

import (
	e "app/events"
	"fmt"
	"log"
	"time"
)

func DefaultUptimeHandler() UptimeHandler {
	return UptimeHandler{
		func(msg string) {
			log.Printf("Uptime is %s", msg)
		},
	}
}

type UptimeHandler struct {
	handle func(string)
}

func (h UptimeHandler) Supports(evtName string) bool {
	return evtName == e.UPTIME
}

func (h UptimeHandler) Handle(evt e.ProcessEvent) error {
	duration, err := time.ParseDuration(fmt.Sprintf("%ss", evt.Data))
	if err == nil {
		h.handle(fmt.Sprintf("%f seconds", duration.Seconds()))
	}

	return nil
}
