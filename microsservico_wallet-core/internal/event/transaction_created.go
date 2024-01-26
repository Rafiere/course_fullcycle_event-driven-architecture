package event

import "time"

type TransactionCreated struct {
	Name    string
	Payload interface{}
}

/* Estamos criando um evento de transação. */
func NewTransactionCreated() *TransactionCreated {

	return &TransactionCreated{
		Name: "TransactionCreated",
	}
}

func (t *TransactionCreated) GetName() string {
	return t.Name
}

func (t *TransactionCreated) GetPayload() interface{} {
	return t.Payload
}

func (t *TransactionCreated) GetDateTime() time.Time {
	return time.Now()
}

func (t *TransactionCreated) SetPayload(payload interface{}) {
	t.Payload = payload
}
