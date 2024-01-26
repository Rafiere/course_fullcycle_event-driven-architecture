package create_account

import (
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ClientGatewayMock struct {
	mock.Mock
}

/* Abaixo, estamos mockando o resultado do "Get", assim, estamos atribuindo o valor desejado. */
func (c *ClientGatewayMock) Get(id string) (*entity.Client, error) {

	args := c.Called(id)

	return args.Get(0).(*entity.Client), args.Error(1)
}

/* Abaixo, estamos mockando o resultado do "Get", assim, estamos atribuindo o valor desejado. */
func (c *ClientGatewayMock) Save(client *entity.Client) error {

	args := c.Called(client)

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

func TestCreateAccountUseCase_Execute(t *testing.T) {

	client, _ := entity.NewClient("John Doe", "j@email.com")
	clientMock := &ClientGatewayMock{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &AccountGatewayMock{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)

	inputDto := CreateAccountInputDTO{
		ClientID: client.ID,
	}

	output, err := uc.Execute(&inputDto)
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	accountMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
