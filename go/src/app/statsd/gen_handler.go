package statsd

import (
	e "app/events"
)

// GenericHandler is a generic message handler compatible with
// the observer.MessageHandler
type GenericHandler struct {
	handle    func(e.ProcessEvent) error
	eventName string
}

// NewGenHandler creates new generic handler with the given handle function and event name
func NewGenHandler(handle func(e.ProcessEvent) error, evtName string) GenericHandler {
	return GenericHandler{
		handle,
		evtName,
	}
}

// Supports tells you if the given eventName is supported by the handler
func (h GenericHandler) Supports(evtName string) bool {
	return evtName == h.eventName
}

// Handle is the event handling function
func (h GenericHandler) Handle(evt e.ProcessEvent) error {
	return h.handle(evt)
}
