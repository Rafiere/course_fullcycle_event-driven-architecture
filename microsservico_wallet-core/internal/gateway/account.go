package gateway

import "github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
