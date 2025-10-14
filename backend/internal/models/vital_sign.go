package models

import (
	"time"

	"gorm.io/gorm"
)

type VitalSign struct {
	ID                     uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	MedicalRecordID        uint           `gorm:"not null;index;unique" json:"medical_record_id" validate:"required"`
	BloodPressureSystolic  *int           `gorm:"" json:"blood_pressure_systolic,omitempty" validate:"omitempty"`
	BloodPressureDiastolic *int           `gorm:"" json:"blood_pressure_diastolic,omitempty" validate:"omitempty"`
	HeartRate              *int           `gorm:"" json:"heart_rate,omitempty" validate:"omitempty"`                         // bpm
	Temperature            *float64       `gorm:"type:decimal(4,2)" json:"temperature,omitempty" validate:"omitempty"`       // Celsius
	RespiratoryRate        *int           `gorm:"" json:"respiratory_rate,omitempty" validate:"omitempty"`                   // per minute
	OxygenSaturation       *float64       `gorm:"type:decimal(5,2)" json:"oxygen_saturation,omitempty" validate:"omitempty"` // percentage
	WeightKg               *float64       `gorm:"type:decimal(5,2)" json:"weight_kg,omitempty" validate:"omitempty"`
	HeightCm               *float64       `gorm:"type:decimal(5,2)" json:"height_cm,omitempty" validate:"omitempty"`
	BMI                    *float64       `gorm:"type:decimal(5,2)" json:"bmi,omitempty" validate:"omitempty"`
	RecordedAt             time.Time      `gorm:"not null;index" json:"recorded_at" validate:"required"`
	CreatedAt              time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	MedicalRecord *MedicalRecord `gorm:"foreignKey:MedicalRecordID;references:ID" json:"medical_record,omitempty"`
}

func (VitalSign) TableName() string {
	return "vital_signs"
}

func (v *VitalSign) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	if v.RecordedAt.IsZero() {
		v.RecordedAt = now
	}
	if v.CreatedAt.IsZero() {
		v.CreatedAt = now
	}
	if v.UpdatedAt.IsZero() {
		v.UpdatedAt = now
	}

	// Calculate BMI if weight and height are provided
	if v.WeightKg != nil && v.HeightCm != nil && *v.HeightCm > 0 {
		heightM := *v.HeightCm / 100
		bmi := *v.WeightKg / (heightM * heightM)
		v.BMI = &bmi
	}

	return nil
}

func (v *VitalSign) CalculateBMI() *float64 {
	if v.WeightKg != nil && v.HeightCm != nil && *v.HeightCm > 0 {
		heightM := *v.HeightCm / 100
		bmi := *v.WeightKg / (heightM * heightM)
		return &bmi
	}
	return nil
}
