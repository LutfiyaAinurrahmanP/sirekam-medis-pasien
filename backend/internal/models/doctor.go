package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         *uint          `gorm:"index" json:"user_id,omitempty" validate:"omitempty"`
	EmployeeID     string         `gorm:"unique;not null;size:20;index" json:"employee_id" validate:"required,max=20"`
	FullName       string         `gorm:"not null;size:100" json:"full_name" validate:"required,max=100"`
	Specialization string         `gorm:"not null;size:100" json:"specialization" validate:"required,max=100"`
	LicenseNumber  string         `gorm:"unique;not null;size:50" json:"license_number" validate:"required,max=50"`
	Phone          string         `gorm:"size:15" json:"phone,omitempty" validate:"omitempty,max=15"`
	Email          string         `gorm:"size:100" json:"email,omitempty" validate:"omitempty,email,max=100"`
	DepartmentID   *uint          `gorm:"index" json:"department_id,omitempty" validate:"omitempty"`
	IsActive       bool           `gorm:"not null;default:true;index" json:"is_active" validate:"required"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	User             *User             `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Department       *Department       `gorm:"foreignKey:DepartmentID;references:ID" json:"department,omitempty"`
	MedicalRecords   []MedicalRecord   `gorm:"foreignKey:DoctorID;references:ID" json:"medical_records,omitempty"`
	Prescriptions    []Prescription    `gorm:"foreignKey:DoctorID;references:ID" json:"prescriptions,omitempty"`
	LabTests         []LabTest         `gorm:"foreignKey:OrderedByDoctorID;references:ID" json:"lab_tests,omitempty"`
	Appointments     []Appointment     `gorm:"foreignKey:DoctorID;references:ID" json:"appointments,omitempty"`
	Hospitalizations []Hospitalization `gorm:"foreignKey:AttendingDoctorID;references:ID" json:"hospitalizations,omitempty"`
}

func (Doctor) TableName() string {
	return "doctors"
}

func (d *Doctor) BeforeCreate(tx *gorm.DB) error {
	if d.IsActive {
		d.IsActive = true
	}

	now := time.Now()
	if d.CreatedAt.IsZero() {
		d.CreatedAt = now
	}
	if d.UpdatedAt.IsZero() {
		d.UpdatedAt = now
	}
	return nil
}

func (d *Doctor) Activate() {
	d.IsActive = true
}

func (d *Doctor) Deactivate() {
	d.IsActive = false
}
