package auth

// RegisterDto ДТО для регистрации пользователя
type RegisterDto struct {
	Email           string `json:"name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

// GetTokenDto ДТО для получения токена
type GetTokenDto struct {
	Email    string `json:"name"`
	Password string `json:"password"`
}

// RefreshTokenDto ДТО для получения токена
type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken"`
}

// Dto ДТО для токена
type Dto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
