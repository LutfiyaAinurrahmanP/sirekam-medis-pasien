package validators

import "time"

type CreateVitalSignRequest struct {
	MedicalRecordID        uint      `json:"medical_record_id" validate:"required"`
	BloodPressureSystolic  *int      `json:"blood_pressure_systolic" validate:"omitempty"`
	BloodPressureDiastolic *int      `json:"blood_pressure_diastolic" validate:"omitempty"`
	HeartRate              *int      `json:"heart_rate" validate:"omitempty"`        // bpm
	Temperature            *float64  `json:"temperature" validate:"omitempty"`       // Celsius
	RespiratoryRate        *int      `json:"respiratory_rate" validate:"omitempty"`  // per minute
	OxygenSaturation       *float64  `json:"oxygen_saturation" validate:"omitempty"` // percentage
	WeightKg               *float64  `json:"weight_kg" validate:"omitempty"`
	HeightCm               *float64  `json:"height_cm" validate:"omitempty"`
	BMI                    *float64  `json:"bmi" validate:"omitempty"`
	RecordedAt             time.Time `json:"recorded_at" validate:"required"`
}

type UpdateVitalSignRequest struct {
	MedicalRecordID        uint      `json:"medical_record_id" validate:"omitempty"`
	BloodPressureSystolic  *int      `json:"blood_pressure_systolic" validate:"omitempty"`
	BloodPressureDiastolic *int      `json:"blood_pressure_diastolic" validate:"omitempty"`
	HeartRate              *int      `json:"heart_rate" validate:"omitempty"`        // bpm
	Temperature            *float64  `json:"temperature" validate:"omitempty"`       // Celsius
	RespiratoryRate        *int      `json:"respiratory_rate" validate:"omitempty"`  // per minute
	OxygenSaturation       *float64  `json:"oxygen_saturation" validate:"omitempty"` // percentage
	WeightKg               *float64  `json:"weight_kg" validate:"omitempty"`
	HeightCm               *float64  `json:"height_cm" validate:"omitempty"`
	BMI                    *float64  `json:"bmi" validate:"omitempty"`
	RecordedAt             time.Time `json:"recorded_at" validate:"omitempty"`
}

type ListVitalSignQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id created_at"`
}

func (q *ListVitalSignQuery) SetVitalSignDefaults() {
	if q.Page < 1 {
		q.Page = 1
	}
	if q.Limit < 1 {
		q.Limit = 10
	}
	if q.Limit > 100 {
		q.Limit = 100
	}
	if q.Sort == "" {
		q.Sort = "desc"
	}
	if q.SortBy == "" {
		q.SortBy = "id"
	}
}

func (q *ListVitalSignQuery) GetVitalSignOffSet() int {
	return (q.Page - 1) * q.Limit
}
