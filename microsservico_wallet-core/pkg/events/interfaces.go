package events

import "time"

/* Essa interface representará um email. */
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{} //Podemos ter diversos payloads nos diversos formatos.
	SetPayload(payload interface{})
}

/* Essa interface representará as operações que são executadas quando um evento é chamado. */

/* É como se fosse um controller, que é executado ao chamarmos uma rota. */
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	/* Para registrarmos o evento, precisamos saber o nome do evento. Quando esse vento
	   acontecer, o "handler" executará o evento. */
	Register(eventName string, handler EventHandlerInterface) error
	Dispatch(event EventInterface) error                          //Os handlers que estão registrados começarão a ser executados.
	Remove(eventName string, handler EventHandlerInterface) error //Estamos removendo o evento.
	Has(eventName string, handler EventHandlerInterface) bool     //Queremos ver se temos um evento com o handler especificado.
	Clear() error                                                 //Esse método limpará o "EventDispatcher", matando todos os eventos que estão registrados ali
}
