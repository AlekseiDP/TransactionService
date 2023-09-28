package filters

import (
	"TransactionService/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

func CheckAuthorized(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		log.Print("Ошибка при получении токена")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Print("Неверный формат токена")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if len(headerParts[1]) == 0 {
		log.Print("Токен пустой")
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.GetJwtConfig().Jwt.Secret), nil
	})

	if err != nil {
		log.Print(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if !token.Valid {
		log.Print("Невалидный токен")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
