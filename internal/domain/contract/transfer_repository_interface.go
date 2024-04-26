package contract

import "github.com/andreluizmicro/desafio-go/internal/domain/entity"

type TransferRepositoryInterface interface {
	Create(transfer entity.Transfer) (int64, error)
}
