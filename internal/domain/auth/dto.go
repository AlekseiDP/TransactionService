package auth

// SignOnDto ДТО для регистрации пользователя
type SignOnDto struct {
	Email           string `json:"name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type Dto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
