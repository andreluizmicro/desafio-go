package contract

import "github.com/andreluizmicro/desafio-go/internal/domain/entity"

type AccountRepositoryInterface interface {
	Create(account entity.Account) (int64, error)
}
