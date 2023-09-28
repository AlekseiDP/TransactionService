package composites

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/api/auth"
	domainAuth "TransactionService/internal/domain/auth"
)

// AuthComposite Структура для регистрации сервиса и хэндлера Auth
type AuthComposite struct {
	Service domainAuth.Service
	Handler api.Handler
}

func NewAuthComposite(p *PostgresComposite) (*AuthComposite, error) {
	authService := domainAuth.NewService(p.DB)
	authHandler := auth.NewHandler(authService)

	return &AuthComposite{
		Service: authService,
		Handler: authHandler,
	}, nil
}
