package filters

import (
	"TransactionService/internal/adapters/handlers"
	"TransactionService/internal/domain/errors/serviceError"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type appHandler func(c *gin.Context) (any, error)

// HandleError Middleware для обработки результата запроса
func HandleError(h appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Модель ответа
		result := handlers.ResultModel{}

		// Обработка ошибок
		var serviceErr *serviceError.ServiceError
		dto, err := h(c)
		if err != nil {
			environment := os.Getenv("ENVIRONMENT")

			if errors.As(err, &serviceErr) {
				serviceErr = err.(*serviceError.ServiceError)
				if environment == "Development" {
					result.ErrorDevelopment = serviceErr.DeveloperMessage
				}

				result.ErrorCode = serviceErr.Code
				result.ErrorMessage = serviceErr.Message

				if serviceErr.Code == "NOT_FOUND" {
					c.JSON(http.StatusNotFound, result)
					return
				} else if serviceErr.Code == "VALIDATION" {
					c.JSON(http.StatusBadRequest, result)
					return
				} else if serviceErr.Code == "REFRESH_TOKEN" || serviceErr.Code == "AUTH" {
					c.JSON(http.StatusUnauthorized, result)
					return
				}

				c.JSON(http.StatusBadRequest, result)
				return
			}

			result.ErrorCode = "UNKNOWN"
			result.ErrorDisplay = "Сервис временно недоступен"
			result.ErrorMessage = "Что-то пошло не так("
			result.ErrorDevelopment = err.Error()
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		result.Data = dto
		c.JSON(http.StatusOK, result)
	}
}
