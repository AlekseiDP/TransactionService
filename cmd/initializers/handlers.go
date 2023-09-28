package initializers

import (
	"TransactionService/internal/adapters/handlers/account"
	"TransactionService/internal/adapters/handlers/auth"
	"TransactionService/internal/adapters/middleware"
	domainAccount "TransactionService/internal/domain/account"
	domainAuth "TransactionService/internal/domain/auth"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterHandlers(engine *gin.Engine) {
	log.Print("Initializing handlers")
	engine.Use(middleware.SetUserId)

	registerAccountHandlers(engine)
	registerAuthHandlers(engine)
}

func registerAccountHandlers(engine *gin.Engine) {
	accountService := domainAccount.NewService(DB)
	accountHandler := account.NewHandler(accountService)
	accountHandler.Register(engine)
}

func registerAuthHandlers(engine *gin.Engine) {
	authService := domainAuth.NewService(DB)
	authHandler := auth.NewHandler(authService)
	authHandler.Register(engine)
}
