package validators

import "time"

type CreateMedicalRecordRequest struct {
	PatientID      uint       `json:"patient_id" validate:"required"`
	DoctorID       uint       `json:"doctor_id" validate:"required"`
	VisitDate      time.Time  `json:"visit_date" validate:"required"`
	VisitTime      string     `json:"visit_time" validate:"required"`
	ChiefComplaint string     `json:"chief_complaint" validate:"required"`
	Symptoms       string     `json:"symptoms" validate:"omitempty"`
	Diagnosis      string     `json:"diagnosis" validate:"required"`
	DiagnosisCode  string     `json:"diagnosis_code" validate:"omitempty,max=20"` // ICD-10
	TreatmentPlan  string     `json:"treatment_plan" validate:"omitempty"`
	Notes          string     `json:"notes" validate:"omitempty"`
	NextVisitDate  *time.Time `json:"next_visit_date" validate:"omitempty"`
	Status         string     `json:"status" validate:"required,oneof=draft completed archived"`
}

type UpdateMedicalRecordRequest struct {
	PatientID      uint       `json:"patient_id" validate:"omitempty"`
	DoctorID       uint       `json:"doctor_id" validate:"omitempty"`
	VisitDate      time.Time  `json:"visit_date" validate:"omitempty"`
	VisitTime      string     `json:"visit_time" validate:"omitempty"`
	ChiefComplaint string     `json:"chief_complaint" validate:"omitempty"`
	Symptoms       string     `json:"symptoms" validate:"omitempty"`
	Diagnosis      string     `json:"diagnosis" validate:"omitempty"`
	DiagnosisCode  string     `json:"diagnosis_code" validate:"omitempty,max=20"` // ICD-10
	TreatmentPlan  string     `json:"treatment_plan" validate:"omitempty"`
	Notes          string     `json:"notes" validate:"omitempty"`
	NextVisitDate  *time.Time `json:"next_visit_date" validate:"omitempty"`
	Status         string     `json:"status" validate:"omitempty,oneof=draft completed archived"`
}

type ListMedicalRecordQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id created_at"`
}

func (q *ListMedicalRecordQuery) SetMedicalRecordDefaults() {
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

func (q *ListMedicalRecordQuery) GetMedicalRecordOffSet() int {
	return (q.Page - 1) * q.Limit
}
