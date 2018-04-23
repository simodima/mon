package observer

import e "app/events"

// MessageHandler is a generic message handler
// that can be injected to the observer
type MessageHandler interface {
	Supports(string) bool
	Handle(e.ProcessEvent) error
}
