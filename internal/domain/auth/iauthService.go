package auth

type Service interface {
	Register(signOnDto *RegisterDto) (*Dto, error)
	GetToken(signInDto *GetTokenDto) (*Dto, error)
	RefreshToken(refreshTokenDto *RefreshTokenDto) (*Dto, error)
}
