package auth

import (
	"TransactionService/internal/config"
	"TransactionService/internal/domain/errors/serviceError"
	"TransactionService/internal/domain/user"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"math/rand"
	"time"
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

func (s service) Register(registerDto *RegisterDto) (*Dto, error) {
	// Получение юзера
	var user1 user.User
	getUser := s.DB.Where("email = ?", registerDto.Email).First(&user1)

	// Если клиента нет то создаем юзера и токены, иначе возвращаем ошибку
	if getUser.Error != nil {
		if getUser.Error.Error() == "record not found" {
			jwtConfig := config.GetJwtConfig()

			// Генерация токенов
			accessToken, err := generateAccessToken(fmt.Sprintf("%v", user1.ID))
			if err != nil {
				return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при создании token", getUser.Error.Error(), "AUTH")
			}

			refreshToken, err := generateRefreshToken()
			if err != nil {
				return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при создании refresh token", getUser.Error.Error(), "AUTH")
			}

			// Создание юзера
			user1.Email = registerDto.Email
			user1.PasswordHash = generatePasswordHash(registerDto.Password)
			user1.RefreshToken = refreshToken
			user1.RefreshExpiresAt = time.Now().Add(jwtConfig.Jwt.RefreshTokenLifetime)

			createUser := s.DB.Create(&user1)
			if createUser.Error != nil {
				return nil, serviceError.NewServiceError(createUser.Error, "Ошибка при создании User", createUser.Error.Error(), "DB")
			}

			result := Dto{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}

			return &result, nil
		}

		return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при получении User", getUser.Error.Error(), "DB")
	}

	return nil, serviceError.NewServiceError(nil, fmt.Sprintf("Пользователь c Email %v уже существует", registerDto.Email), fmt.Sprintf("Пользователь c Email %v уже существует", registerDto.Email), "SIGN_ON")
}

func (s service) GetToken(getTokenDto *GetTokenDto) (*Dto, error) {
	// Получение юзера
	var user1 user.User
	getUser := s.DB.Where("email = ?", getTokenDto.Email).First(&user1)

	// Если клиента нет то возвращаем ошибку, иначе создаем токены
	if getUser.Error != nil {
		if getUser.Error.Error() == "record not found" {
			return nil, serviceError.NewServiceError(getUser.Error, fmt.Sprintf("User c Email %v не существует", getTokenDto.Email), getUser.Error.Error(), "NOT_FOUND")
		}

		return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при получении User", getUser.Error.Error(), "DB")
	}

	// Проверка пароля
	passwordHash := generatePasswordHash(getTokenDto.Password)
	if passwordHash != user1.PasswordHash {
		return nil, serviceError.NewServiceError(nil, "Неверный пароль", "Неверный пароль", "AUTH")
	}

	// Генерация токенов
	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, serviceError.NewServiceError(nil, "Ошибка при создании refresh token", "Ошибка при создании refresh token", "AUTH")
	}

	accessToken, err := generateAccessToken(fmt.Sprintf("%v", user1.ID))
	if err != nil {
		return nil, serviceError.NewServiceError(nil, "Ошибка при создании token", "Ошибка при создании refresh token", "AUTH")
	}

	result := Dto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &result, nil
}

func (s service) RefreshToken(refreshTokenDto *RefreshTokenDto) (*Dto, error) {
	// Получение юзера
	var user1 user.User
	getUser := s.DB.Where("refreshToken = ?", refreshTokenDto.RefreshToken).First(&user1)

	// Если клиента нет то возвращаем ошибку, иначе создаем токены
	if getUser.Error != nil {
		if getUser.Error.Error() == "record not found" {
			return nil, serviceError.NewServiceError(getUser.Error, "User не найден", getUser.Error.Error(), "NOT_FOUND")
		}

		return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при получении User", getUser.Error.Error(), "DB")
	}

	// Проверка жизни refresh token
	if user1.RefreshExpiresAt.After(time.Now()) {
		return nil, serviceError.NewServiceError(nil, "Refresh token истек", "Refresh token истек", "REFRESH_TOKEN")
	}

	// Генерация токенов
	jwtConfig := config.GetJwtConfig()

	accessToken, err := generateAccessToken(fmt.Sprintf("%v", user1.ID))
	if err != nil {
		return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при создании token", getUser.Error.Error(), "AUTH")
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return nil, serviceError.NewServiceError(getUser.Error, "Ошибка при создании refresh token", getUser.Error.Error(), "AUTH")
	}

	updateUser := s.DB.Save(&user.User{ID: 1, RefreshToken: refreshToken, RefreshExpiresAt: time.Now().Add(jwtConfig.Jwt.RefreshTokenLifetime)})
	if updateUser.Error != nil {
		return nil, serviceError.NewServiceError(updateUser.Error, "Ошибка при обновлении User", updateUser.Error.Error(), "DB")
	}

	result := Dto{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &result, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash)
}

func generateAccessToken(userId string) (string, error) {
	jwtConfig := config.GetJwtConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    jwtConfig.Jwt.Issuer,
		Audience:  jwt.ClaimStrings{jwtConfig.Jwt.Audience},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.Jwt.AccessTokenLifetime)),
		Subject:   userId,
	})

	return token.SignedString([]byte(jwtConfig.Jwt.Secret))
}

func generateRefreshToken() (string, error) {
	bytes := make([]byte, 32)

	source := rand.NewSource(time.Now().Unix())
	random := rand.New(source)

	_, err := random.Read(bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}
