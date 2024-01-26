package gateway

import "github.com.br/Rafiere/course_fullcycle_arquitetura-baseada-em-microsservicos/microsservico_wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
