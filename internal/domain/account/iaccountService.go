package account

import (
	"TransactionService/internal/adapters/handlers"
)

type Service interface {
	Create(createAccountDto *CreateAccountDto) (*Dto, error)
	ListPage(pageIndex, pageSize int) (*handlers.PageableModel, error)
}
