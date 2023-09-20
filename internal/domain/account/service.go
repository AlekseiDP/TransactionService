package account

import (
	apiAccount "TransactionService/internal/adapters/api/account"
	"github.com/devfeel/mapper"
	"gorm.io/gorm"
)

type service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) apiAccount.Service {
	return &service{
		DB: db,
	}
}

func (s *service) Create(createAccountDto *apiAccount.CreateAccountDto) (*apiAccount.Dto, error) {
	item := Account{Owner: createAccountDto.Owner, Balance: createAccountDto.Balance, Currency: createAccountDto.Currency}
	result := s.DB.Create(&item)
	if result.Error != nil {
		return nil, result.Error
	}

	var dto apiAccount.Dto
	if err := mapper.AutoMapper(&item, &dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (s *service) ListPage(pageIndex, pageSize int) (*apiAccount.PageDto, error) {
	var items []Account
	var count int64
	result1 := s.DB.Limit(pageSize).Offset(pageIndex * pageSize).Find(&items)
	if result1.Error != nil {
		return nil, result1.Error
	}

	result2 := s.DB.Model(&items).Count(&count)
	if result2.Error != nil {
		return nil, result2.Error
	}

	var dto []apiAccount.ShortDto
	if err := mapper.MapperSlice(&items, &dto); err != nil {
		return nil, err
	}

	pageDto := apiAccount.PageDto{
		Items:            dto,
		CurrentPageIndex: pageIndex,
		TotalCount:       count,
	}

	return &pageDto, nil
}
