package validators

import "time"

type CreateBillingRequest struct {
	PatientID       uint       `json:"patient_id" validate:"required"`
	MedicalRecordID *uint      `json:"medical_record_id" validate:"omitempty"`
	InvoiceNumber   string     `json:"invoice_number" validate:"required,max=50"`
	InvoiceDate     time.Time  `json:"invoice_date" validate:"required"`
	TotalAmount     float64    `json:"total_amount" validate:"required,min=0"`
	PaidAmount      float64    `json:"paid_amount" validate:"omitempty"`
	DiscountAmount  *float64   `json:"discount_amount" validate:"omitempty"`
	PaymentStatus   string     `json:"payment_status" validate:"required,oneof=unpaid partial paid cancelled"`
	PaymentMethod   string     `json:"payment_method" validate:"omitempty,oneof=cash debit_card credit_card insurance transfer other"`
	PaymentDate     *time.Time `json:"payment_date" validate:"omitempty"`
	Notes           string     `json:"notes" validate:"omitempty"`
}

type UpdateBillingRequest struct {
	PatientID       uint       `json:"patient_id" validate:"omitempty"`
	MedicalRecordID *uint      `json:"medical_record_id" validate:"omitempty"`
	InvoiceNumber   string     `json:"invoice_number" validate:"omitempty,max=50"`
	InvoiceDate     time.Time  `json:"invoice_date" validate:"omitempty"`
	TotalAmount     float64    `json:"total_amount" validate:"omitempty,min=0"`
	PaidAmount      float64    `json:"paid_amount" validate:"omitempty"`
	DiscountAmount  *float64   `json:"discount_amount" validate:"omitempty"`
	PaymentStatus   string     `json:"payment_status" validate:"omitempty,oneof=unpaid partial paid cancelled"`
	PaymentMethod   string     `json:"payment_method" validate:"omitempty,oneof=cash debit_card credit_card insurance transfer other"`
	PaymentDate     *time.Time `json:"payment_date" validate:"omitempty"`
	Notes           string     `json:"notes" validate:"omitempty"`
}

type ListBillingQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id invoice_date payment_status payment_method payment_date created_at"`
}

func (q *ListBillingQuery) SetBillingDefaults() {
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

func (q *ListBillingQuery) GetBillingOffSet() int {
	return (q.Page - 1) * q.Limit
}
