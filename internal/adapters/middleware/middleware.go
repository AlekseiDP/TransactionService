package middleware

import (
	"TransactionService/internal/adapters/api"
	"TransactionService/internal/domain/errors/serviceError"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type appHandler func(c *gin.Context) (any, error)

// Middleware Middleware для обработки результата запроса
func Middleware(h appHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Модель ответа
		result := api.ResultModel{}

		// Обработка ошибок
		var serviceErr *serviceError.ServiceError
		dto, err := h(c)
		if err != nil {
			if errors.As(err, &serviceErr) {
				serviceErr = err.(*serviceError.ServiceError)
				if serviceErr.Code == "NOT_FOUND" {
					result.ErrorCode = serviceErr.Code
					result.ErrorMessage = serviceErr.DeveloperMessage
					c.JSON(http.StatusNotFound, result)
					return
				} else if serviceErr.Code == "VALIDATION" {
					result.ErrorCode = serviceErr.Code
					result.ErrorMessage = serviceErr.DeveloperMessage
					c.JSON(http.StatusBadRequest, result)
					return
				}

				result.ErrorCode = serviceErr.Code
				result.ErrorMessage = err.Error()
				c.JSON(http.StatusInternalServerError, result)
				return
			}

			result.ErrorCode = "UNKNOWN"
			result.ErrorDisplay = "Сервис временно недоступен"
			result.ErrorMessage = "Something went wrong"
			c.JSON(http.StatusInternalServerError, result)
			return
		}

		result.Data = dto
		c.JSON(http.StatusOK, result)
	}
}
