package account

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/middleware"
	"TransactionService/internal/domain/account"
	"TransactionService/internal/domain/errors/serviceError"
	"github.com/gin-gonic/gin"
)

const (
	createAccountUrl   = "/accounts"
	listAccountPageUrl = "/accounts/paginated"
)

// handler структура для обработки Http запросов
type handler struct {
	accountService account.Service
}

func NewHandler(service account.Service) api.Handler {
	return &handler{accountService: service}
}

func (h *handler) Register(engine *gin.Engine) {
	engine.POST(createAccountUrl, middleware.ErrorMiddleware(h.CreateAccount))
	engine.GET(listAccountPageUrl, middleware.ErrorMiddleware(h.ListAccountPage))
}

func (h *handler) CreateAccount(c *gin.Context) (any, error) {
	// Получение входных данных
	createAccountDto := account.CreateAccountDto{}
	if err := c.ShouldBindJSON(&createAccountDto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при получении входных данных", err.Error(), "BIND")
	}

	// Вызов сервиса
	dto, err := h.accountService.Create(&createAccountDto)
	if err != nil {
		return nil, err
	}

	return dto, err
}

func (h *handler) ListAccountPage(c *gin.Context) (any, error) {
	// Получение входных данных
	var params struct {
		PageIndex int `form:"pageIndex"`
		PageSize  int `form:"pageSize"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при получении входных данных", err.Error(), "BIND")
	}

	// Вызов сервиса
	dto, err := h.accountService.ListPage(params.PageIndex, params.PageSize)
	if err != nil {
		return nil, err
	}

	return dto, err
}
