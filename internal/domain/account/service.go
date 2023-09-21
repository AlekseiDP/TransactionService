package account

import (
	"TransactionService/internal/adapters/api"
	apiAccount "TransactionService/internal/adapters/api/account"
	"TransactionService/internal/domain/errors/serviceError"
	"github.com/devfeel/mapper"
	"gorm.io/gorm"
)

// service Структура для обработки бизнес логики
type service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) apiAccount.Service {
	return &service{
		DB: db,
	}
}

// Create Функция создания записи Account
func (s *service) Create(createAccountDto *apiAccount.CreateAccountDto) (*apiAccount.Dto, error) {
	item := Account{Owner: createAccountDto.Owner, Balance: createAccountDto.Balance, Currency: createAccountDto.Currency}
	result := s.DB.Create(&item)
	if result.Error != nil {
		return nil, serviceError.NewServiceError(result.Error, "Ошибка при создании Account", result.Error.Error(), "DB")
	}

	var dto apiAccount.Dto
	if err := mapper.AutoMapper(&item, &dto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при маппинге Account", err.Error(), "MAP")
	}

	return &dto, nil
}

// ListPage Функция получения постраничного списка Account
func (s *service) ListPage(pageIndex, pageSize int) (*api.PageableModel, error) {
	var items []Account
	var count int64
	result1 := s.DB.Limit(pageSize).Offset(pageIndex * pageSize).Find(&items)
	if result1.Error != nil {
		return nil, serviceError.NewServiceError(result1.Error, "Ошибка при получении постраничного списка Account", result1.Error.Error(), "DB")
	}

	result2 := s.DB.Model(&items).Count(&count)
	if result2.Error != nil {
		return nil, serviceError.NewServiceError(result2.Error, "Ошибка при получении количества Account", result2.Error.Error(), "DB")
	}

	var dto []apiAccount.ShortDto
	if err := mapper.MapperSlice(&items, &dto); err != nil {
		return nil, serviceError.NewServiceError(err, "Ошибка при маппинге Account", err.Error(), "MAP")
	}

	pageDto := api.PageableModel{
		Items:            dto,
		CurrentPageIndex: pageIndex,
		TotalCount:       count,
	}

	return &pageDto, nil
}
