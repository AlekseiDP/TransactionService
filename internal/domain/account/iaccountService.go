package account

import (
	"TransactionService/internal/adapters/api"
)

type Service interface {
	Create(createAccountDto *CreateAccountDto) (*Dto, error)
	ListPage(pageIndex, pageSize int) (*api.PageableModel, error)
}
