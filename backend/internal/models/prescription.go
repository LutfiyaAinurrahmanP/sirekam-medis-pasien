package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	PrescriptionStatusPending   = "pending"
	PrescriptionStatusDispensed = "dispensed"
	PrescriptionStatusCancelled = "cancelled"
)

type Prescription struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	MedicalRecordID  uint           `gorm:"not null;index" json:"medical_record_id" validate:"required"`
	DoctorID         uint           `gorm:"not null;index" json:"doctor_id" validate:"required"`
	PrescriptionDate time.Time      `gorm:"not null;type:date;index" json:"prescription_date" validate:"required"`
	Notes            string         `gorm:"type:text" json:"notes,omitempty"`
	Status           string         `gorm:"type:enum('pending','dispensed','cancelled');not null;default:'pending';index" json:"status"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	MedicalRecord *MedicalRecord     `gorm:"foreignKey:MedicalRecordID;references:ID" json:"medical_record,omitempty"`
	Doctor        *Doctor            `gorm:"foreignKey:DoctorID;references:ID" json:"doctor,omitempty"`
	Items         []PrescriptionItem `gorm:"foreignKey:PrescriptionID;references:ID" json:"items,omitempty"`
}

func (Prescription) TableName() string {
	return "prescriptions"
}

func (p *Prescription) BeforeCreate(tx *gorm.DB) error {
	if p.Status == "" {
		p.Status = PrescriptionStatusPending
	}

	now := time.Now()
	if p.PrescriptionDate.IsZero() {
		p.PrescriptionDate = now
	}
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = now
	}
	return nil
}

func (p *Prescription) MarkAsDispensed() {
	p.Status = PrescriptionStatusDispensed
}

func (p *Prescription) Cancel() {
	p.Status = PrescriptionStatusCancelled
}

func ValidatePrescriptionStatus(status string) bool {
	return status == PrescriptionStatusPending ||
		status == PrescriptionStatusDispensed ||
		status == PrescriptionStatusCancelled
}

func GetAvailablePrescriptionStatuses() []string {
	return []string{
		PrescriptionStatusPending,
		PrescriptionStatusDispensed,
		PrescriptionStatusCancelled,
	}
}