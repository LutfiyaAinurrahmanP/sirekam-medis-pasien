package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	AppointmentStatusScheduled  = "scheduled"
	AppointmentStatusConfirmed  = "confirmed"
	AppointmentStatusInProgress = "in_progress"
	AppointmentStatusCompleted  = "completed"
	AppointmentStatusCancelled  = "cancelled"
	AppointmentStatusNoShow     = "no_show"
)

type Appointment struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID       uint           `gorm:"not null;index" json:"patient_id" validate:"required"`
	DoctorID        uint           `gorm:"not null;index" json:"doctor_id" validate:"required"`
	AppointmentDate time.Time      `gorm:"not null;type:date;index" json:"appointment_date" validate:"required"`
	AppointmentTime string         `gorm:"not null;type:time" json:"appointment_time" validate:"required"` // HH:MM:SS
	DurationMinutes int            `gorm:"not null" json:"duration_minutes" validate:"required,min=1"`
	Reason          string         `gorm:"type:text" json:"reason,omitempty" validate:"omitempty"`
	Status          string         `gorm:"type:enum('scheduled','confirmed','in_progress','completed','cancelled','no_show');not null;default:'scheduled';index" json:"status" validate:"required,oneof=scheduled confirmed in_progress completed cancelled no_show"`
	Notes           string         `gorm:"type:text" json:"notes,omitempty" validate:"omitempty"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient *Patient `gorm:"foreignKey:PatientID;references:ID" json:"patient,omitempty"`
	Doctor  *Doctor  `gorm:"foreignKey:DoctorID;references:ID" json:"doctor,omitempty"`
}

func (Appointment) TableName() string {
	return "appointments"
}

func (a *Appointment) BeforeCreate(tx *gorm.DB) error {
	if a.Status == "" {
		a.Status = AppointmentStatusScheduled
	}

	now := time.Now()
	if a.CreatedAt.IsZero() {
		a.CreatedAt = now
	}
	if a.UpdatedAt.IsZero() {
		a.UpdatedAt = now
	}
	return nil
}

func (a *Appointment) Confirm() {
	a.Status = AppointmentStatusConfirmed
}

func (a *Appointment) StartProgress() {
	a.Status = AppointmentStatusInProgress
}

func (a *Appointment) Complete() {
	a.Status = AppointmentStatusCompleted
}

func (a *Appointment) Cancel() {
	a.Status = AppointmentStatusCancelled
}

func (a *Appointment) MarkAsNoShow() {
	a.Status = AppointmentStatusNoShow
}

func ValidateAppointmentStatus(status string) bool {
	validStatuses := GetAvailableAppointmentStatuses()
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}

func GetAvailableAppointmentStatuses() []string {
	return []string{
		AppointmentStatusScheduled,
		AppointmentStatusConfirmed,
		AppointmentStatusInProgress,
		AppointmentStatusCompleted,
		AppointmentStatusCancelled,
		AppointmentStatusNoShow,
	}
}
