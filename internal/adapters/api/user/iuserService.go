package user

type Service interface {
	Create(createUserDto *CreateUserDto) (*Dto, error)
	GetByEmail(email string) (*Dto, error)
	GetByRefreshToken(refreshToken string) (*Dto, error)
}
