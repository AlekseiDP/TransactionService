package composites

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/api/account"
	domainAccount "TransactionService/internal/domain/account"
)

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
