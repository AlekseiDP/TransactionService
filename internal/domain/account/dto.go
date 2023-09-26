package account

import "time"

// CreateAccountDto ДТО для создания Account
type CreateAccountDto struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

// Dto ДТО для Account
type Dto struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ShortDto ДТО для Account
type ShortDto struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
}
