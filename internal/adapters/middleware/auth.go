package middleware

import (
	"TransactionService/internal/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func SetUserId(c *gin.Context) {
	header := c.GetHeader("Authorization")
	headerParts := strings.Split(header, "")

	if header != "" {
		token, err := jwt.Parse(headerParts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.GetJwtConfig().Jwt.Secret), nil
		})

		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("UserId", claims["sub"])
			}
		}
	}
}
