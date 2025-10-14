package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	MedicalRecordStatusDraft     = "draft"
	MedicalRecordStatusCompleted = "completed"
	MedicalRecordStatusArchived  = "archived"
)

type MedicalRecord struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	PatientID     uint           `gorm:"not null;index" json:"patient_id" validate:"required"`
	DoctorID      uint           `gorm:"not null;index" json:"doctor_id" validate:"required"`
	VisitDate     time.Time      `gorm:"not null;type:date;index" json:"visit_date" validate:"required"`
	VisitTime     string         `gorm:"not null;type:time" json:"visit_time" validate:"required"`
	ChiefComplaint string        `gorm:"not null;type:text" json:"chief_complaint" validate:"required"`
	Symptoms      string         `gorm:"type:text" json:"symptoms,omitempty"`
	Diagnosis     string         `gorm:"not null;type:text" json:"diagnosis" validate:"required"`
	DiagnosisCode string         `gorm:"size:20" json:"diagnosis_code,omitempty"` // ICD-10
	TreatmentPlan string         `gorm:"type:text" json:"treatment_plan,omitempty"`
	Notes         string         `gorm:"type:text" json:"notes,omitempty"`
	NextVisitDate *time.Time     `gorm:"type:date" json:"next_visit_date,omitempty"`
	Status        string         `gorm:"type:enum('draft','completed','archived');not null;default:'draft';index" json:"status"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient       *Patient       `gorm:"foreignKey:PatientID;references:ID" json:"patient,omitempty"`
	Doctor        *Doctor        `gorm:"foreignKey:DoctorID;references:ID" json:"doctor,omitempty"`
	VitalSigns    *VitalSign     `gorm:"foreignKey:MedicalRecordID;references:ID" json:"vital_signs,omitempty"`
	Prescriptions []Prescription `gorm:"foreignKey:MedicalRecordID;references:ID" json:"prescriptions,omitempty"`
	LabTests      []LabTest      `gorm:"foreignKey:MedicalRecordID;references:ID" json:"lab_tests,omitempty"`
}

func (MedicalRecord) TableName() string {
	return "medical_records"
}

func (m *MedicalRecord) BeforeCreate(tx *gorm.DB) error {
	if m.Status == "" {
		m.Status = MedicalRecordStatusDraft
	}

	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

func (m *MedicalRecord) MarkAsCompleted() {
	m.Status = MedicalRecordStatusCompleted
}

func (m *MedicalRecord) Archive() {
	m.Status = MedicalRecordStatusArchived
}

func ValidateMedicalRecordStatus(status string) bool {
	return status == MedicalRecordStatusDraft ||
		status == MedicalRecordStatusCompleted ||
		status == MedicalRecordStatusArchived
}

func GetAvailableMedicalRecordStatuses() []string {
	return []string{
		MedicalRecordStatusDraft,
		MedicalRecordStatusCompleted,
		MedicalRecordStatusArchived,
	}
}