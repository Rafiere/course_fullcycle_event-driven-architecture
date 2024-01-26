package entity

import (
	"time"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) *Account {

	if client == nil {
		return nil
	}

	account := &Account{
		ID:        client.ID,
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) Credit(value float64) {
	a.Balance += value
	a.UpdatedAt = time.Now()
}

func (a *Account) Debit(value float64) {
	a.Balance -= value
	a.UpdatedAt = time.Now()
}
