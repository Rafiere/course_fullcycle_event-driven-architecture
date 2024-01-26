package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "johndoe@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "johndoe@email.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe Updated", "email@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Updated", client.Name)
	assert.Equal(t, "email@email.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "email@email.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@email.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
