package composites

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/api/account"
	domainAccount "TransactionService/internal/domain/account"
)

// AccountComposite Структура для регистрации сервиса и хэндлера Account
type AccountComposite struct {
	Service account.Service
	Handler api.Handler
}

func NewAccountComposite(p *PostgresComposite) (*AccountComposite, error) {
	accountService := domainAccount.NewService(p.DB)
	accountHandler := account.NewHandler(accountService)

	return &AccountComposite{
		Service: accountService,
		Handler: accountHandler,
	}, nil
}
