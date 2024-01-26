package create_client

import (
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"
	"github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/gateway"
	"time"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutputDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

/* Isso é como se fosse um construtor na orientação a objetos. */
func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (c *CreateClientUseCase) Execute(input *CreateClientInputDTO) (*CreateClientOutputDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}

	err = c.ClientGateway.Save(client)
	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDTO{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil
}
