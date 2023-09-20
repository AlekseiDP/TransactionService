package account

type Service interface {
	Create(createAccountDto *CreateAccountDto) (*Dto, error)
}
