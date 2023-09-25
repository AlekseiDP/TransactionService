package user

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/middleware"
	"TransactionService/internal/domain/errors/serviceError"
	"github.com/gin-gonic/gin"
)

const (
	createUserUrl = "/users"
	getUserUrl    = "/users"
)

// handler структура для обработки Http запросов
type handler struct {
	userService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(engine *gin.Engine) {
	engine.POST(createUserUrl, middleware.Middleware(h.CreateUser))
	engine.GET(getUserUrl, middleware.Middleware(h.GetUser))
}

func (h *handler) CreateUser(c *gin.Context) (any, error) {
	// Получение входных данных
	createAccountDto := CreateUserDto{}
	if err := c.ShouldBindJSON(&createAccountDto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при получении входных данных", err.Error(), "BIND")
	}

	// Вызов сервиса
	dto, err := h.userService.Create(&createAccountDto)
	if err != nil {
		return nil, err
	}

	return dto, err
}

func (h *handler) GetUser(c *gin.Context) (any, error) {
	// Получение входных данных
	var params struct {
		Email string `form:"email"`
	}

	if err := c.ShouldBindQuery(&params); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при получении входных данных", err.Error(), "BIND")
	}

	// Вызов сервиса
	dto, err := h.userService.GetByEmail(params.Email)
	if err != nil {
		return nil, err
	}

	return dto, err
}
