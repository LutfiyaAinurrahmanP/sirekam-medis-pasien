package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	LabTestStatusOrdered         = "ordered"
	LabTestStatusSampleCollected = "sample_collected"
	LabTestStatusInProgress      = "in_progress"
	LabTestStatusCompleted       = "completed"
	LabTestStatusCancelled       = "cancelled"
)

type LabTest struct {
	ID                   uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	MedicalRecordID      uint           `gorm:"not null;index" json:"medical_record_id" validate:"required"`
	TestTypeID           uint           `gorm:"not null;index" json:"test_type_id" validate:"required"`
	OrderedByDoctorID    uint           `gorm:"not null;index" json:"ordered_by_doctor_id" validate:"required"`
	OrderDate            time.Time      `gorm:"not null;type:date;index" json:"order_date" validate:"required"`
	SampleCollectionDate *time.Time     `gorm:"type:datetime" json:"sample_collection_date,omitempty" validate:"omitempty"`
	ResultDate           *time.Time     `gorm:"type:datetime" json:"result_date,omitempty" validate:"omitempty"`
	ResultValue          string         `gorm:"type:text" json:"result_value,omitempty" validate:"omitempty"`
	ResultUnit           string         `gorm:"size:50" json:"result_unit,omitempty" validate:"omitempty,max=50"`
	ReferenceRange       string         `gorm:"size:100" json:"reference_range,omitempty" validate:"omitempty,max=100"`
	Status               string         `gorm:"type:enum('ordered','sample_collected','in_progress','completed','cancelled');not null;default:'ordered';index" json:"status" validate:"required,oneof=ordered sample_collected in_progress completed cancelled"`
	Notes                string         `gorm:"type:text" json:"notes,omitempty" validate:"omitempty"`
	CreatedAt            time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	MedicalRecord   *MedicalRecord `gorm:"foreignKey:MedicalRecordID;references:ID" json:"medical_record,omitempty"`
	TestType        *TestType      `gorm:"foreignKey:TestTypeID;references:ID" json:"test_type,omitempty"`
	OrderedByDoctor *Doctor        `gorm:"foreignKey:OrderedByDoctorID;references:ID" json:"ordered_by_doctor,omitempty"`
}

func (LabTest) TableName() string {
	return "lab_tests"
}

func (l *LabTest) BeforeCreate(tx *gorm.DB) error {
	if l.Status == "" {
		l.Status = LabTestStatusOrdered
	}

	now := time.Now()
	if l.OrderDate.IsZero() {
		l.OrderDate = now
	}
	if l.CreatedAt.IsZero() {
		l.CreatedAt = now
	}
	if l.UpdatedAt.IsZero() {
		l.UpdatedAt = now
	}
	return nil
}

func (l *LabTest) MarkSampleCollected() {
	l.Status = LabTestStatusSampleCollected
	now := time.Now()
	l.SampleCollectionDate = &now
}

func (l *LabTest) MarkInProgress() {
	l.Status = LabTestStatusInProgress
}

func (l *LabTest) MarkCompleted() {
	l.Status = LabTestStatusCompleted
	now := time.Now()
	l.ResultDate = &now
}

func (l *LabTest) Cancel() {
	l.Status = LabTestStatusCancelled
}

func ValidateLabTestStatus(status string) bool {
	validStatuses := GetAvailableLabTestStatuses()
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}

func GetAvailableLabTestStatuses() []string {
	return []string{
		LabTestStatusOrdered,
		LabTestStatusSampleCollected,
		LabTestStatusInProgress,
		LabTestStatusCompleted,
		LabTestStatusCancelled,
	}
}
