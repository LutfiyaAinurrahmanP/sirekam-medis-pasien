package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null;size:100;index" json:"name" validate:"required"`
	Code          string         `gorm:"unique;not null;size:20;index" json:"code" validate:"required"`
	Description   string         `gorm:"type:text" json:"description,omitempty"`
	FloorLocation string         `gorm:"size:50" json:"floor_location,omitempty"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Doctors []Doctor `gorm:"foreignKey:DepartmentID;references:ID" json:"doctors,omitempty"`
	Rooms   []Room   `gorm:"foreignKey:DepartmentID;references:ID" json:"rooms,omitempty"`
}

func (Department) TableName() string {
	return "departments"
}

func (d *Department) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	if d.CreatedAt.IsZero() {
		d.CreatedAt = now
	}
	if d.UpdatedAt.IsZero() {
		d.UpdatedAt = now
	}
	return nil
}