package validators

type CreatePrescriptionItemRequest struct {
	PrescriptionID uint   `json:"prescription_id" validate:"required"`
	MedicineID     uint   `json:"medicine_id" validate:"required"`
	Dosage         string `json:"dosage" validate:"required,max=100"`    // e.g., "500mg"
	Frequency      string `json:"frequency" validate:"required,max=100"` // e.g., "3x sehari"
	DurationDays   int    `json:"duration_days" validate:"required,min=1"`
	Quantity       int    `json:"quantity" validate:"required,min=1"`
	Instructions   string `json:"instructions" validate:"omitempty"`
}

type UpdatePrescriptionItemRequest struct {
	PrescriptionID uint   `json:"prescription_id" validate:"omitempty"`
	MedicineID     uint   `json:"medicine_id" validate:"omitempty"`
	Dosage         string `json:"dosage" validate:"omitempty,max=100"`    // e.g., "500mg"
	Frequency      string `json:"frequency" validate:"omitempty,max=100"` // e.g., "3x sehari"
	DurationDays   int    `json:"duration_days" validate:"omitempty,min=1"`
	Quantity       int    `json:"quantity" validate:"omitempty,min=1"`
	Instructions   string `json:"instructions" validate:"omitempty"`
}

type ListPrescriptionItemQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id created_at"`
}

func (q *ListPrescriptionItemQuery) SetPrescriptionItemDefaults() {
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

func (q *ListPrescriptionItemQuery) GetPrescriptionItemOffSet() int {
	return (q.Page - 1) * q.Limit
}