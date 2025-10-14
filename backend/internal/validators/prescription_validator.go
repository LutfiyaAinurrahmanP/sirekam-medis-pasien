package validators

import "time"

type CreatePrescriptionRequest struct {
	MedicalRecordID  uint      `json:"medical_record_id" validate:"required"`
	DoctorID         uint      `json:"doctor_id" validate:"required"`
	PrescriptionDate time.Time `json:"prescription_date" validate:"required"`
	Notes            string    `json:"notes" validate:"omitempty"`
	Status           string    `json:"status" validate:"required,oneof=pending dispensed cancelled"`
}

type UpdatePrescriptionRequest struct {
	MedicalRecordID  uint      `json:"medical_record_id" validate:"omitempty"`
	DoctorID         uint      `json:"doctor_id" validate:"omitempty"`
	PrescriptionDate time.Time `json:"prescription_date" validate:"omitempty"`
	Notes            string    `json:"notes" validate:"omitempty"`
	Status           string    `json:"status" validate:"omitempty,oneof=pending dispensed cancelled"`
}

type ListPrescriptionQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id created_at"`
}

func (q *ListPrescriptionQuery) SetPrescriptionDefaults() {
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

func (q *ListPrescriptionQuery) GetPrescriptionOffSet() int {
	return (q.Page - 1) * q.Limit
}
