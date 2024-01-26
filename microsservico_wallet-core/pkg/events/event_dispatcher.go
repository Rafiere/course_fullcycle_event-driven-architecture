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

func (ed *EventDispatcher) Clear() {
	/* Estamos zerando os handlers. */
	ed.handlers = make(map[string][]EventHandlerInterface)
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) {
	if _, ok := ed.handlers[event.GetName()]; ok {

		/* Para cada um dos "handler", executaremos o método "handle". */
		for _, handler := range ed.handlers[event.GetName()] {
			handler.Handle(event)
		}
	}
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) {
	if _, ok := ed.handlers[eventName]; ok { //Se tivermos um elemento registrado, passaremos para o "for".
		for i, h := range ed.handlers[eventName] {
			if h == handler { //Se o handler for igual ao handler que queremos remover, faremos a remoção.
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...) //Removendo o handler.
				return
			}
		}
	}
}
