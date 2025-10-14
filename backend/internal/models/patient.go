package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	GenderMale   = "male"
	GenderFemale = "female"
	GenderOther  = "other"
)

type Patient struct {
	ID                    uint           `gorm:"primaryKey" json:"id"`
	UserID                *uint          `gorm:"index" json:"user_id,omitempty" validate:"omitempty"`
	PatientCode           string         `gorm:"unique;not null;size:20;index" json:"patient_code" validate:"required,max=20"`
	FullName              string         `gorm:"not null;size:100" json:"full_name" validate:"required,max=100"`
	DateOfBirth           time.Time      `gorm:"not null;type:date" json:"date_of_birth" validate:"required"`
	Gender                string         `gorm:"type:enum('male','female','other');not null" json:"gender" validate:"required,oneof=male female other"`
	BloodType             string         `gorm:"size:5" json:"blood_type,omitempty" validate:"omitempty,max=5"`
	Phone                 string         `gorm:"size:15" json:"phone,omitempty" validate:"omitempty,max=15"`
	Email                 string         `gorm:"size:100" json:"email,omitempty" validate:"omitempty,email,max=100"`
	Address               string         `gorm:"type:text" json:"address,omitempty" validate:"omitempty"`
	EmergencyContactName  string         `gorm:"size:100" json:"emergency_contact_name,omitempty" validate:"omitempty,max=100"`
	EmergencyContactPhone string         `gorm:"size:15" json:"emergency_contact_phone,omitempty" validate:"omitempty,max=15"`
	InsuranceNumber       string         `gorm:"size:50" json:"insurance_number,omitempty" validate:"omitempty,max=50"`
	InsuranceProvider     string         `gorm:"size:100" json:"insurance_provider,omitempty" validate:"omitempty,max=100"`
	Allergies             string         `gorm:"type:text" json:"allergies,omitempty" validate:"omitempty"`
	CreatedAt             time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	User             *User              `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	MedicalRecords   []MedicalRecord    `gorm:"foreignKey:PatientID;references:ID" json:"medical_records,omitempty"`
	Appointments     []Appointment      `gorm:"foreignKey:PatientID;references:ID" json:"appointments,omitempty"`
	Hospitalizations []Hospitalization  `gorm:"foreignKey:PatientID;references:ID" json:"hospitalizations,omitempty"`
	Billings         []Billing          `gorm:"foreignKey:PatientID;references:ID" json:"billings,omitempty"`
}

func (Patient) TableName() string {
	return "patients"
}

func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = now
	}
	return nil
}

func (p *Patient) GetAge() int {
	now := time.Now()
	age := now.Year() - p.DateOfBirth.Year()
	if now.YearDay() < p.DateOfBirth.YearDay() {
		age--
	}
	return age
}

func ValidateGender(gender string) bool {
	return gender == GenderMale || gender == GenderFemale || gender == GenderOther
}

func GetAvailableGenders() []string {
	return []string{GenderMale, GenderFemale, GenderOther}
}