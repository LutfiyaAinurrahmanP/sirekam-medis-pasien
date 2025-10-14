package validators

type CreateDepartmentRequest struct {
	Name          string `json:"name" validate:"required,max=100"`
	Code          string `json:"code" validate:"required,max=20"`
	Description   string `json:"description" validate:"omitempty"`
	FloorLocation string `json:"floor_location" validate:"omitempty,max=50"`
}

type UpdateDepartmentRequest struct {
	Name          string `json:"name" validate:"omitempty,max=100"`
	Code          string `json:"code" validate:"omitempty,max=20"`
	Description   string `json:"description" validate:"omitempty"`
	FloorLocation string `json:"floor_location" validate:"omitempty,max=50"`
}

type ListDepartmentQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id created_at"`
}

func (q *ListDepartmentQuery) SetDepartmentDefaults() {
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

func (q *ListDepartmentQuery) GetDepartmentOffSet() int {
	return (q.Page - 1) * q.Limit
}
