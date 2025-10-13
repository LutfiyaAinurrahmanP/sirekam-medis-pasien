package models

import "time"

type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"type:varchar(512);not null;index" json:"token"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	ExpiresAt time.Time `gorm:"not null;index" json:"expires_at"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (TokenBlacklist) TableName() string {
	return "token_blacklist"
}

func (t *TokenBlacklist) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}
