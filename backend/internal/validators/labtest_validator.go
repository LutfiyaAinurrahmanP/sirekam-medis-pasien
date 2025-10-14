package validators

import "time"

type CreateLabTestRequest struct {
	MedicalRecordID      uint       `json:"medical_record_id" validate:"required"`
	TestTypeID           uint       `json:"test_type_id" validate:"required"`
	OrderedByDoctorID    uint       `json:"ordered_by_doctor_id" validate:"required"`
	OrderDate            time.Time  `json:"order_date" validate:"required"`
	SampleCollectionDate *time.Time `json:"sample_collection_date" validate:"omitempty"`
	ResultDate           *time.Time `json:"result_date" validate:"omitempty"`
	ResultValue          string     `json:"result_value" validate:"omitempty"`
	ResultUnit           string     `json:"result_unit" validate:"omitempty,max=50"`
	ReferenceRange       string     `json:"reference_range" validate:"omitempty,max=100"`
	Status               string     `json:"status" validate:"required,oneof=ordered sample_collected in_progress completed cancelled"`
	Notes                string     `json:"notes" validate:"omitempty"`
}

type UpdateLabTestRequest struct {
	MedicalRecordID      uint       `json:"medical_record_id" validate:"omitempty"`
	TestTypeID           uint       `json:"test_type_id" validate:"omitempty"`
	OrderedByDoctorID    uint       `json:"ordered_by_doctor_id" validate:"omitempty"`
	OrderDate            time.Time  `json:"order_date" validate:"omitempty"`
	SampleCollectionDate *time.Time `json:"sample_collection_date" validate:"omitempty"`
	ResultDate           *time.Time `json:"result_date" validate:"omitempty"`
	ResultValue          string     `json:"result_value" validate:"omitempty"`
	ResultUnit           string     `json:"result_unit" validate:"omitempty,max=50"`
	ReferenceRange       string     `json:"reference_range" validate:"omitempty,max=100"`
	Status               string     `json:"status" validate:"omitempty,oneof=ordered sample_collected in_progress completed cancelled"`
	Notes                string     `json:"notes" validate:"omitempty"`
}

type ListLabTestQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id order_date created_at"`
}

func (q *ListLabTestQuery) SetLabTestDefaults() {
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

func (q *ListLabTestQuery) GetLabTestOffSet() int {
	return (q.Page - 1) * q.Limit
}