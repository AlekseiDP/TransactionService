package auth

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/adapters/middleware"
	"TransactionService/internal/domain/auth"
	"TransactionService/internal/domain/errors/serviceError"
	"github.com/gin-gonic/gin"
)

const (
	signOnUrl   = "/connect/sign-on"
	getTokenUrl = "/connect/token"
)

// handler структура для обработки Http запросов
type handler struct {
	authService auth.Service
}

func NewHandler(service auth.Service) api.Handler {
	return &handler{authService: service}
}

func (h *handler) Register(engine *gin.Engine) {
	engine.POST(signOnUrl, middleware.ErrorMiddleware(h.SignOn))
	engine.POST(getTokenUrl, middleware.ErrorMiddleware(h.GetToken))
}

func (h *handler) SignOn(c *gin.Context) (any, error) {
	// Получение входных данных
	registerDto := auth.RegisterDto{}
	registerDto.Email = c.PostForm("email")
	registerDto.Password = c.PostForm("password")
	registerDto.PasswordConfirm = c.PostForm("passwordConfirm")

	// Вызов сервиса
	dto, err := h.authService.Register(&registerDto)

	return dto, err
}

func (h *handler) GetToken(c *gin.Context) (any, error) {
	// Получение входных данных
	grantType := c.PostForm("grant_type")
	if grantType == "password" {
		getTokenDto := auth.GetTokenDto{}
		getTokenDto.Email = c.PostForm("email")
		getTokenDto.Password = c.PostForm("password")

		// Вызов сервиса
		dto, err := h.authService.GetToken(&getTokenDto)

		return dto, err
	} else if grantType == "refresh_token" {
		refreshTokenDto := auth.RefreshTokenDto{}
		refreshTokenDto.RefreshToken = c.PostForm("refreshToken")

		// Вызов сервиса
		dto, err := h.authService.RefreshToken(&refreshTokenDto)

		return dto, err
	}

	return nil, serviceError.NewServiceError(nil, "Некорректный grant_types", "Некорректный grant_types", "AUTH")
}
