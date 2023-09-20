package account

import "time"

type CreateAccountDto struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type Dto struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ShortDto struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
}

type PageDto struct {
	Items            any   `json:"items"`
	CurrentPageIndex int   `json:"currentPageIndex"`
	TotalCount       int64 `json:"totalCount"`
}
