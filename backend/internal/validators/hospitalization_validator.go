package validators

import "time"

type CreateHospitalizationRequest struct {
	PatientID         uint       `json:"patient_id" validate:"required"`
	MedicalRecordID   uint       `json:"medical_record_id" validate:"required"`
	AdmissionDate     time.Time  `json:"admission_date" validate:"required"`
	DischargeDate     *time.Time `json:"discharge_date" validate:"omitempty"`
	RoomID            uint       `json:"room_id" validate:"required"`
	AttendingDoctorID uint       `json:"attending_doctor_id" validate:"required"`
	AdmissionReason   string     `json:"admission_reason" validate:"required"`
	DischargeSummary  string     `json:"discharge_summary" validate:"omitempty"`
	Status            string     `json:"status" validate:"required,oneof=admitted discharged transferred"`
}

type UpdateHospitalizationRequest struct {
	PatientID         uint       `json:"patient_id" validate:"omitempty"`
	MedicalRecordID   uint       `json:"medical_record_id" validate:"omitempty"`
	AdmissionDate     time.Time  `json:"admission_date" validate:"omitempty"`
	DischargeDate     *time.Time `json:"discharge_date" validate:"omitempty"`
	RoomID            uint       `json:"room_id" validate:"omitempty"`
	AttendingDoctorID uint       `json:"attending_doctor_id" validate:"omitempty"`
	AdmissionReason   string     `json:"admission_reason" validate:"omitempty"`
	DischargeSummary  string     `json:"discharge_summary" validate:"omitempty"`
	Status            string     `json:"status" validate:"omitempty,oneof=admitted discharged transferred"`
}

type ListHospitalizationQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id created_at"`
}

func (q *ListHospitalizationQuery) SetHospitalizationDefaults() {
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

func (q *ListHospitalizationQuery) GetHospitalizationOffSet() int {
	return (q.Page - 1) * q.Limit
}
