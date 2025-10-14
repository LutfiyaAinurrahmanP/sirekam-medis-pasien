package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	HospitalizationStatusAdmitted    = "admitted"
	HospitalizationStatusDischarged  = "discharged"
	HospitalizationStatusTransferred = "transferred"
)

type Hospitalization struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	PatientID         uint           `gorm:"not null;index" json:"patient_id" validate:"required"`
	MedicalRecordID   uint           `gorm:"not null;index" json:"medical_record_id" validate:"required"`
	AdmissionDate     time.Time      `gorm:"not null;type:datetime;index" json:"admission_date" validate:"required"`
	DischargeDate     *time.Time     `gorm:"type:datetime" json:"discharge_date,omitempty" validate:"omitempty"`
	RoomID            uint           `gorm:"not null;index" json:"room_id" validate:"required"`
	AttendingDoctorID uint           `gorm:"not null;index" json:"attending_doctor_id" validate:"required"`
	AdmissionReason   string         `gorm:"not null;type:text" json:"admission_reason" validate:"required"`
	DischargeSummary  string         `gorm:"type:text" json:"discharge_summary,omitempty" validate:"omitempty"`
	Status            string         `gorm:"type:enum('admitted','discharged','transferred');not null;default:'admitted';index" json:"status" validate:"required"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient         *Patient       `gorm:"foreignKey:PatientID;references:ID" json:"patient,omitempty"`
	MedicalRecord   *MedicalRecord `gorm:"foreignKey:MedicalRecordID;references:ID" json:"medical_record,omitempty"`
	Room            *Room          `gorm:"foreignKey:RoomID;references:ID" json:"room,omitempty"`
	AttendingDoctor *Doctor        `gorm:"foreignKey:AttendingDoctorID;references:ID" json:"attending_doctor,omitempty"`
}

func (Hospitalization) TableName() string {
	return "hospitalizations"
}

func (h *Hospitalization) BeforeCreate(tx *gorm.DB) error {
	if h.Status == "" {
		h.Status = HospitalizationStatusAdmitted
	}

	now := time.Now()
	if h.AdmissionDate.IsZero() {
		h.AdmissionDate = now
	}
	if h.CreatedAt.IsZero() {
		h.CreatedAt = now
	}
	if h.UpdatedAt.IsZero() {
		h.UpdatedAt = now
	}
	return nil
}

func (h *Hospitalization) Discharge() {
	h.Status = HospitalizationStatusDischarged
	now := time.Now()
	h.DischargeDate = &now
}

func (h *Hospitalization) Transfer() {
	h.Status = HospitalizationStatusTransferred
}

func (h *Hospitalization) GetDurationDays() int {
	endDate := time.Now()
	if h.DischargeDate != nil {
		endDate = *h.DischargeDate
	}
	duration := endDate.Sub(h.AdmissionDate)
	days := int(duration.Hours() / 24)
	if days < 1 {
		return 1 // minimum 1 day
	}
	return days
}

func ValidateHospitalizationStatus(status string) bool {
	return status == HospitalizationStatusAdmitted ||
		status == HospitalizationStatusDischarged ||
		status == HospitalizationStatusTransferred
}

func GetAvailableHospitalizationStatuses() []string {
	return []string{
		HospitalizationStatusAdmitted,
		HospitalizationStatusDischarged,
		HospitalizationStatusTransferred,
	}
}