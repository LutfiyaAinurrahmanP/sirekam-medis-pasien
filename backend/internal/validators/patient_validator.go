package validators

import "time"

type CreatePatientRequest struct {
	UserID                *uint     `json:"user_id" validate:"omitempty"`
	PatientCode           string    `json:"patient_code" validate:"required,max=20"`
	FullName              string    `json:"full_name" validate:"required,max=100"`
	DateOfBirth           time.Time `json:"date_of_birth" validate:"required"`
	Gender                string    `json:"gender" validate:"required,oneof=male female other"`
	BloodType             string    `json:"blood_type" validate:"omitempty,max=5"`
	Phone                 string    `json:"phone" validate:"omitempty,max=15"`
	Email                 string    `json:"email" validate:"omitempty,email,max=100"`
	Address               string    `json:"address" validate:"omitempty"`
	EmergencyContactName  string    `json:"emergency_contact_name" validate:"omitempty,max=100"`
	EmergencyContactPhone string    `json:"emergency_contact_phone" validate:"omitempty,max=15"`
	InsuranceNumber       string    `json:"insurance_number" validate:"omitempty,max=50"`
	InsuranceProvider     string    `json:"insurance_provider" validate:"omitempty,max=100"`
	Allergies             string    `json:"allergies" validate:"omitempty"`
}

type UpdatePatientRequest struct {
	UserID                *uint     `json:"user_id" validate:"omitempty"`
	PatientCode           string    `json:"patient_code" validate:"omitempty,max=20"`
	FullName              string    `json:"full_name" validate:"omitempty,max=100"`
	DateOfBirth           time.Time `json:"date_of_birth" validate:"omitempty"`
	Gender                string    `json:"gender" validate:"omitempty,oneof=male female other"`
	BloodType             string    `json:"blood_type" validate:"omitempty,max=5"`
	Phone                 string    `json:"phone" validate:"omitempty,max=15"`
	Email                 string    `json:"email" validate:"omitempty,email,max=100"`
	Address               string    `json:"address" validate:"omitempty"`
	EmergencyContactName  string    `json:"emergency_contact_name" validate:"omitempty,max=100"`
	EmergencyContactPhone string    `json:"emergency_contact_phone" validate:"omitempty,max=15"`
	InsuranceNumber       string    `json:"insurance_number" validate:"omitempty,max=50"`
	InsuranceProvider     string    `json:"insurance_provider" validate:"omitempty,max=100"`
	Allergies             string    `json:"allergies" validate:"omitempty"`
}

type ListPatientQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id gender date_of_birth created_at"`
}

func (q *ListPatientQuery) SetPatientDefaults() {
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

func (q *ListPatientQuery) GetPatientOffSet() int {
	return (q.Page - 1) * q.Limit
}
