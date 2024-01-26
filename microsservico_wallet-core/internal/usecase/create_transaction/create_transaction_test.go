package create_transaction

import (
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (t *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := t.Called(transaction)
	return args.Error(0)
}

/* Abaixo, temos os mocks de "account" */
type AccountGatewayMock struct {
	mock.Mock
}

func (a *AccountGatewayMock) Save(account *entity.Account) error {

	args := a.Called(account)

	return args.Error(0)
}

func (a *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {

	args := a.Called(id)

	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {

	client1, _ := entity.NewClient("client1", "client1@email.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)

	client2, _ := entity.NewClient("client2", "client2@email.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindByID", account1.ID).Return(account1, nil)
	mockAccount.On("FindByID", account2.ID).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount)

	output, err := uc.Execute(&inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindByID", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
