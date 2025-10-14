package validators

type CreateTestTypeRequest struct {
	Name        string   `json:"name" validate:"required,max=200"`
	Code        string   `json:"code" validate:"required,max=50"`
	Category    string   `json:"category" validate:"omitempty,max=100"` // Hematologi, Kimia Darah, etc
	Description string   `json:"description" validate:"omitempty"`
	Price       *float64 `json:"price" validate:"omitempty"`
	IsActive    bool     `json:"is_active" validate:"required"`
}

type UpdateTestTypeRequest struct {
	Name        string   `json:"name" validate:"omitempty,max=200"`
	Code        string   `json:"code" validate:"omitempty,max=50"`
	Category    string   `json:"category" validate:"omitempty,max=100"` // Hematologi, Kimia Darah, etc
	Description string   `json:"description" validate:"omitempty"`
	Price       *float64 `json:"price" validate:"omitempty"`
	IsActive    bool     `json:"is_active" validate:"omitempty"`
}

type ListTestTypeQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id category created_at"`
}

func (q *ListTestTypeQuery) SetTestTypeDefaults() {
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

func (q *ListTestTypeQuery) GetTestTypeOffSet() int {
	return (q.Page - 1) * q.Limit
}