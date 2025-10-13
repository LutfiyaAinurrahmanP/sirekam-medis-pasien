package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null;size:50" json:"username" validate:"required,min=3,max=50"`
	Email     string         `gorm:"unique;not null;size:100" json:"email" validate:"required,email"`
	Phone     string         `gorm:"unique;not null;size:15" json:"phone" validate:"required,min=10,max=15"`
	Password  string         `gorm:"not null;size:255" json:"password" validate:"required,min=8"`
	Role      string         `gorm:"type:varchar(20);not null;default:'user';index" json:"role"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = RoleUser
	}

	now := time.Now()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = now
	}
	return nil
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

func (u *User) IsUser() bool {
	return u.Role == RoleUser
}

func ValidateRole(role string) bool {
	return role == RoleUser || role == RoleAdmin
}

func GetAvailableRoles() []string {
	return []string{RoleUser, RoleAdmin}
}
