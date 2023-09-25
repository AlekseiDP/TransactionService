package composites

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/api/user"
	domainUser "TransactionService/internal/domain/user"
)

// UserComposite Структура для регистрации сервиса и хэндлера User
type UserComposite struct {
	Service user.Service
	Handler api.Handler
}

func NewUserComposite(p *PostgresComposite) (*UserComposite, error) {
	userService := domainUser.NewService(p.DB)
	userHandler := user.NewHandler(userService)

	return &UserComposite{
		Service: userService,
		Handler: userHandler,
	}, nil
}
