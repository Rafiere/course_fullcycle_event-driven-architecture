package events

import "errors"

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

// Um evento poderá ter vários handlers.
type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface //Um evento poderá ter vários handlers
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

/* Um evento terá vários handlers. */
func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {

	if _, ok = e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	/* Se o handler nunca foi adicionado ao evento, abaixo, faremos essa adição. */
	e.handlers[eventName] = append(e.handlers[eventName], handler)
	return nil
}
