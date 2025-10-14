package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoleUser       = "user"
	RoleAdmin      = "admin"
	RoleDoctor     = "doctor"
	RoleNurse      = "nurse"
	RolePharmacist = "pharmacist"
	RoleLabTech    = "lab_technician"
	RoleReceptionist = "receptionist"
	RoleCashier    = "cashier"
	RoleHeadDept   = "head_department"
	RoleDirector   = "director"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique;not null;size:50" json:"username" validate:"required,min=3,max=50"`
	Email     string         `gorm:"unique;not null;size:100" json:"email" validate:"required,email"`
	Phone     string         `gorm:"unique;not null;size:15" json:"phone" validate:"required,min=10,max=15"`
	Password  string         `gorm:"not null;size:255" json:"password" validate:"required,min=8"`
	Role      string         `gorm:"type:varchar(20);not null;default:'user';index" json:"role" validate:"required"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient *Patient `gorm:"foreignKey:UserID;references:ID" json:"patient,omitempty"`
	Doctor  *Doctor  `gorm:"foreignKey:UserID;references:ID" json:"doctor,omitempty"`
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

func (u *User) IsDoctor() bool {
	return u.Role == RoleDoctor
}

func (u *User) IsNurse() bool {
	return u.Role == RoleNurse
}

func (u *User) IsPharmacist() bool {
	return u.Role == RolePharmacist
}

func (u *User) IsLabTech() bool {
	return u.Role == RoleLabTech
}

func (u *User) IsReceptionist() bool {
	return u.Role == RoleReceptionist
}

func (u *User) IsCashier() bool {
	return u.Role == RoleCashier
}

func (u *User) IsHeadDept() bool {
	return u.Role == RoleHeadDept
}

func (u *User) IsDirector() bool {
	return u.Role == RoleDirector
}

func ValidateRole(role string) bool {
	validRoles := GetAvailableRoles()
	for _, r := range validRoles {
		if r == role {
			return true
		}
	}
	return false
}

func GetAvailableRoles() []string {
	return []string{
		RoleUser,
		RoleAdmin,
		RoleDoctor,
		RoleNurse,
		RolePharmacist,
		RoleLabTech,
		RoleReceptionist,
		RoleCashier,
		RoleHeadDept,
		RoleDirector,
	}
}