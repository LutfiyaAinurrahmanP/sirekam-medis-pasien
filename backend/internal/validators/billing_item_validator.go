package validators

type CreateBillingItemRequest struct {
	BillingID       uint    `json:"billing_id" validate:"required"`
	ItemType        string  `json:"item_type" validate:"required,oneof=consultation medicine lab_test procedure room other"`
	ItemDescription string  `json:"item_description" validate:"required,max=255"`
	Quantity        int     `json:"quantity" validate:"required,min=1"`
	UnitPrice       float64 `json:"unit_price" validate:"required,min=0"`
	TotalPrice      float64 `json:"total_price" validate:"omitempty,min=0"`
}

type UpdateBillingItemRequest struct {
	BillingID       uint    `json:"billing_id" validate:"omitempty"`
	ItemType        string  `json:"item_type" validate:"omitempty,oneof=consultation medicine lab_test procedure room other"`
	ItemDescription string  `json:"item_description" validate:"omitempty,max=255"`
	Quantity        int     `json:"quantity" validate:"omitempty,min=1"`
	UnitPrice       float64 `json:"unit_price" validate:"omitempty,min=0"`
	TotalPrice      float64 `json:"total_price" validate:"omitempty,min=0"`
}

type ListBillingItemQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id item_type created_at"`
}

func (q *ListBillingItemQuery) SetBillingItemDefaults() {
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

func (q *ListBillingItemQuery) GetBillingItemOffSet() int {
	return (q.Page - 1) * q.Limit
}
