package account

import (
	apiAccount "TransactionService/internal/adapters/api/account"
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
	account := Account{Owner: createAccountDto.Owner, Balance: createAccountDto.Balance, Currency: createAccountDto.Currency}
	result := s.DB.Create(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	dto := apiAccount.Dto{ID: account.ID, Owner: account.Owner, Balance: account.Balance, Currency: account.Currency, CreatedAt: account.CreatedAt, UpdatedAt: account.UpdatedAt}

	return &dto, nil
}
