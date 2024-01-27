package create_client

import (
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

/* Testaremos o caso de uso de forma de unidade. Não testaremos o ato de
salvar no banco de dados. */

type ClientGatewayMock struct {
	mock.Mock
}

/* Abaixo, estamos mockando o resultado do "FindByID", assim, estamos atribuindo o valor desejado. */
func (c *ClientGatewayMock) Get(id string) (*entity.Client, error) {

	args := c.Called(id)

	return args.Get(0).(*entity.Client), args.Error(1)
}

/* Abaixo, estamos mockando o resultado do "FindByID", assim, estamos atribuindo o valor desejado. */
func (c *ClientGatewayMock) Save(client *entity.Client) error {

	args := c.Called(client)

	return args.Error(0)
}

/*
	Com esse teste unitário mockando a chamada ao banco de dados, sabemos que, se algum erro

ocorrer, o erro está na chamada ao banco de dados, e não no caso de uso.
*/
func TestCreateClientUseCase_Execute(t *testing.T) {

	m := &ClientGatewayMock{}

	/* Quando passarmos um parâmetro para o "save", será retornado "nil". */
	m.On("Save", mock.Anything).Return(nil)

	uc := NewCreateClientUseCase(m)

	output, err := uc.Execute(&CreateClientInputDTO{
		Name:  "John Doe",
		Email: "j@j.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "John Doe", output.Name)
	assert.Equal(t, "j@j.com", output.Email)
	m.AssertExpectations(t)             // Verifica se o mock foi chamado.
	m.AssertNumberOfCalls(t, "Save", 1) // Verifica se o mock foi chamado apenas uma vez.
}
