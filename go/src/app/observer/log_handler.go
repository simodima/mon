package observer

import (
	e "app/events"
	"fmt"
	"log"
)

func DefaultLogger() Logger {
	return Logger{
		func(msg string) {
			log.Printf("LOG %s", msg)
		},
	}
}

type Logger struct {
	log func(string)
}

func (l Logger) Supports(evtName string) bool {
	return evtName == "LOG"
}

func (l Logger) Handle(evt e.ProcessEvent) error {
	l.log(fmt.Sprint(evt.Message, evt.Data))

	return nil
}
