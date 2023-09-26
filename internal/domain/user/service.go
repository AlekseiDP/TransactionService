package user

import (
	"TransactionService/internal/domain/errors/serviceError"
	"fmt"
	"github.com/devfeel/mapper"
	"gorm.io/gorm"
)

// service Структура для обработки бизнес логики
type service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		DB: db,
	}
}

// Create Функция создания записи User
func (s *service) Create(createUserDto *CreateUserDto) (*Dto, error) {
	var item User
	if err := mapper.AutoMapper(&createUserDto, &item); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при маппинге createUserDto", err.Error(), "MAP")
	}

	result := s.DB.Create(&item)
	if result.Error != nil {
		return nil, serviceError.NewServiceError(result.Error, "Ошибка при создании User", result.Error.Error(), "DB")
	}

	var dto Dto
	if err := mapper.AutoMapper(&item, &dto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при маппинге User", err.Error(), "MAP")
	}

	return &dto, nil
}

// GetByEmail Функция получения записи User по Email
func (s *service) GetByEmail(email string) (*Dto, error) {
	var item User
	result := s.DB.Where("email = ?", email).First(&item)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, serviceError.NewServiceError(result.Error, fmt.Sprintf("User c Email %v не существует", email), result.Error.Error(), "NOT_FOUND")
		}

		return nil, serviceError.NewServiceError(result.Error, "Ошибка при получении User", result.Error.Error(), "DB")
	}

	var dto Dto
	if err := mapper.AutoMapper(&item, &dto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при маппинге User", err.Error(), "MAP")
	}

	return &dto, nil
}

// GetByRefreshToken Функция получения записи User по RefreshToken
func (s *service) GetByRefreshToken(refreshToken string) (*Dto, error) {
	var item User
	result := s.DB.Where("refreshToken = ?", refreshToken).First(&item)
	if result.Error != nil {
		return nil, serviceError.NewServiceError(result.Error, "Ошибка при получении User", result.Error.Error(), "DB")
	}

	var dto Dto
	if err := mapper.AutoMapper(&item, &dto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при маппинге User", err.Error(), "MAP")
	}

	return &dto, nil
}
