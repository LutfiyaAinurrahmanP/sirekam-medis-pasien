package models

import (
	"time"

	"gorm.io/gorm"
)

type TestType struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;size:200;index" json:"name" validate:"required,max=200"`
	Code        string         `gorm:"unique;not null;size:50;index" json:"code" validate:"required,max=50"`
	Category    string         `gorm:"size:100" json:"category,omitempty" validate:"omitempty,max=100"` // Hematologi, Kimia Darah, etc
	Description string         `gorm:"type:text" json:"description,omitempty" validate:"omitempty"`
	Price       *float64       `gorm:"type:decimal(10,2)" json:"price,omitempty" validate:"omitempty"`
	IsActive    bool           `gorm:"not null;default:true;index" json:"is_active" validate:"required"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	LabTests []LabTest `gorm:"foreignKey:TestTypeID;references:ID" json:"lab_tests,omitempty"`
}

func (TestType) TableName() string {
	return "test_types"
}

func (t *TestType) BeforeCreate(tx *gorm.DB) error {
	if t.IsActive {
		t.IsActive = true
	}

	now := time.Now()
	if t.CreatedAt.IsZero() {
		t.CreatedAt = now
	}
	if t.UpdatedAt.IsZero() {
		t.UpdatedAt = now
	}
	return nil
}

func (t *TestType) Activate() {
	t.IsActive = true
}

func (t *TestType) Deactivate() {
	t.IsActive = false
}