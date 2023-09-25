package user

import "time"

// User модель сущности для БД
type User struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"not null" json:"name"`
	PasswordHash string    `gorm:"not null" json:"passwordHash"`
	RefreshToken string    `json:"refreshToken"`
	LockoutEndAt time.Time `json:"lockoutEndAt"`
	RegisteredAt time.Time `gorm:"autoCreateTime" json:"registeredAt"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
