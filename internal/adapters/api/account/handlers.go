package account

import (
	"TransactionService/internal/adapters"
	"TransactionService/internal/adapters/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	createAccountUrl = "/account"
)

type handler struct {
	accountService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{accountService: service}
}

func (h *handler) Register(engine *gin.Engine) {
	engine.POST(createAccountUrl, h.CreateAccount)
}

func (h *handler) CreateAccount(c *gin.Context) {
	// Модель ответа
	result := adapters.ApiResult{}

	// Получение входных данных
	createAccountDto := CreateAccountDto{}
	if err := c.ShouldBindJSON(&createAccountDto); err != nil {
		return
	}

	// Вызов сервиса
	dto, err := h.accountService.Create(&createAccountDto)

	// Формирование ответа
	if err != nil {
		result.ErrorCode = "UNKNOWN"
		result.ErrorDisplay = "Something went wrong"
		result.ErrorMessage = err.Error()
		c.JSON(http.StatusBadRequest, result)
		return
	}
	result.Data = dto
	c.JSON(http.StatusOK, result)
}
