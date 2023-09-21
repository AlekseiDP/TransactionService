package account

import (
	"time"
)

// Account модель сущности для БД
type Account struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Owner     string    `gorm:"not null" json:"owner"`
	Balance   int64     `gorm:"not null" json:"balance"`
	Currency  string    `gorm:"not null" json:"currency"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
