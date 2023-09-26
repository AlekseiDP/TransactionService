package auth

import (
	"TransactionService/internal/config"
	"TransactionService/internal/domain/errors/serviceError"
	"TransactionService/internal/domain/user"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// service Структура для обработки бизнес логики
type service struct {
	DB *gorm.DB
}

func NewService(DB *gorm.DB) Service {
	return &service{
		DB: DB,
	}
}

func (s service) SignOn(signOnDto *SignOnDto) (*Dto, error) {
	var user1 user.User
	result := s.DB.Where("email = ?", signOnDto.Email).First(&user1)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			user1.Email = signOnDto.Email
			user1.PasswordHash = generatePasswordHash(signOnDto.Password)

			result := s.DB.Create(&user1)
			if result.Error != nil {
				return nil, serviceError.NewServiceError(result.Error, "Ошибка при создании User", result.Error.Error(), "DB")
			}
		}

		return nil, serviceError.NewServiceError(result.Error, "Ошибка при получении User", result.Error.Error(), "DB")
	}

	return nil, serviceError.NewServiceError(nil, fmt.Sprintf("Пользователь c Email %v уже существует", signOnDto.Email), fmt.Sprintf("Пользователь c Email %v уже существует", signOnDto.Email), "SIGN_ON")
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func generateAccessToken(user *user.User) *jwt.Token {
	jwtConfig := config.GetJwtConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:   jwtConfig.Jwt.Issuer,
		Audience: jwt.ClaimStrings{jwtConfig.Jwt.Audience},
	})

	return token
}
