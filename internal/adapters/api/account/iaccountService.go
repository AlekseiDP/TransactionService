package account

type Service interface {
	Create(createAccountDto *CreateAccountDto) (*Dto, error)
	ListPage(pageIndex, pageSize int) (*PageDto, error)
}
