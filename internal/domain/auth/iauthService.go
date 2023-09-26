package auth

type Service interface {
	SignOn(signOnDto *SignOnDto) (*Dto, error)
}
