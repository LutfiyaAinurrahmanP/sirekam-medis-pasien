package validators

type CreateMedicineRequest struct {
	Name          string   `json:"name" validate:"required,max=200"`
	GenericName   string   `json:"generic_name" validate:"omitempty,max=200"`
	BrandName     string   `json:"brand_name" validate:"omitempty,max=200"`
	Type          string   `json:"type" validate:"required,oneof=tablet capsule syrup injection ointment other"`
	Strength      string   `json:"strength" validate:"omitempty,max=50"` // e.g., "500mg"
	Manufacturer  string   `json:"manufacturer" validate:"omitempty,max=100"`
	Unit          string   `json:"unit" validate:"omitempty,max=20"` // tablet, ml, mg
	StockQuantity int      `json:"stock_quantity" validate:"required"`
	Price         *float64 `json:"price" validate:"omitempty"`
	IsActive      bool     `json:"is_active" validate:"required"`
}

type UpdateMedicineRequest struct {
	Name          string   `json:"name" validate:"omitempty,max=200"`
	GenericName   string   `json:"generic_name" validate:"omitempty,max=200"`
	BrandName     string   `json:"brand_name" validate:"omitempty,max=200"`
	Type          string   `json:"type" validate:"omitempty,oneof=tablet capsule syrup injection ointment other"`
	Strength      string   `json:"strength" validate:"omitempty,max=50"` // e.g., "500mg"
	Manufacturer  string   `json:"manufacturer" validate:"omitempty,max=100"`
	Unit          string   `json:"unit" validate:"omitempty,max=20"` // tablet, ml, mg
	StockQuantity int      `json:"stock_quantity" validate:"omitempty"`
	Price         *float64 `json:"price" validate:"omitempty"`
	IsActive      bool     `json:"is_active" validate:"omitempty"`
}

type ListMedicineQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id type stock_quantity created_at"`
}

func (q *ListMedicineQuery) SetMedicineDefaults() {
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

func (q *ListMedicineQuery) GetMedicineOffSet() int {
	return (q.Page - 1) * q.Limit
}
