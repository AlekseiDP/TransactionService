package account

import (
	"TransactionService/internal/adapters/filters"
	"TransactionService/internal/adapters/handlers"
	"TransactionService/internal/domain/account"
	"TransactionService/internal/domain/errors/serviceError"
	"github.com/gin-gonic/gin"
)

// handler структура для обработки Http запросов
type handler struct {
	accountService account.Service
}

const (
	createAccountUrl = "/api/v1/accounts"
	listAccountsUrl  = "/api/v1/accounts/paged"
)

func NewHandler(service account.Service) handlers.Handler {
	return &handler{accountService: service}
}

func (h *handler) Register(engine *gin.Engine) {
	authorized := engine.Group("/", filters.CheckAuthorized)
	{
		authorized.POST(createAccountUrl, filters.HandleError(h.CreateAccount))
		authorized.GET(listAccountsUrl, filters.HandleError(h.ListAccountPage))
	}
}

func (h *handler) CreateAccount(c *gin.Context) (any, error) {
	// Получение входных данных
	createAccountDto := account.CreateAccountDto{}
	if err := c.ShouldBindJSON(&createAccountDto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при получении входных данных", err.Error(), "BIND")
	}

	// Вызов сервиса
	dto, err := h.accountService.Create(&createAccountDto)

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

	return dto, err
}
