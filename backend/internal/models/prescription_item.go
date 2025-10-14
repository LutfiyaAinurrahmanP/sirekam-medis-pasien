package models

import (
	"time"

	"gorm.io/gorm"
)

type PrescriptionItem struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	PrescriptionID uint           `gorm:"not null;index" json:"prescription_id" validate:"required"`
	MedicineID     uint           `gorm:"not null;index" json:"medicine_id" validate:"required"`
	Dosage         string         `gorm:"not null;size:100" json:"dosage" validate:"required,max=100"` // e.g., "500mg"
	Frequency      string         `gorm:"not null;size:100" json:"frequency" validate:"required,max=100"` // e.g., "3x sehari"
	DurationDays   int            `gorm:"not null" json:"duration_days" validate:"required,min=1"`
	Quantity       int            `gorm:"not null" json:"quantity" validate:"required,min=1"`
	Instructions   string         `gorm:"type:text" json:"instructions,omitempty" validate:"omitempty"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Prescription *Prescription `gorm:"foreignKey:PrescriptionID;references:ID" json:"prescription,omitempty"`
	Medicine     *Medicine     `gorm:"foreignKey:MedicineID;references:ID" json:"medicine,omitempty"`
}

func (PrescriptionItem) TableName() string {
	return "prescription_items"
}

func (p *PrescriptionItem) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = now
	}
	return nil
}