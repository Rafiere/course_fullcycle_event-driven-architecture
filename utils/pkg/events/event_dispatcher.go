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

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}

	//Se um handler não foi adicionado anteriormente nesse evento, faremos a adição.
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	return nil
}
