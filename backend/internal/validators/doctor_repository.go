package validators

type CreateDoctorRequest struct {
	UserID         *uint  `json:"user_id" validate:"omitempty"`
	EmployeeID     string `json:"employee_id" validate:"required,max=20"`
	FullName       string `json:"full_name" validate:"required,max=100"`
	Specialization string `json:"specialization" validate:"required,max=100"`
	LicenseNumber  string `json:"license_number" validate:"required,max=50"`
	Phone          string `json:"phone" validate:"omitempty,max=15"`
	Email          string `json:"email" validate:"omitempty,email,max=100"`
	DepartmentID   *uint  `json:"department_id" validate:"omitempty"`
	IsActive       bool   `json:"is_active" validate:"required"`
}

type UpdateDoctorRequest struct {
	UserID         *uint  `json:"user_id" validate:"omitempty"`
	EmployeeID     string `json:"employee_id" validate:"omitempty,max=20"`
	FullName       string `json:"full_name" validate:"omitempty,max=100"`
	Specialization string `json:"specialization" validate:"omitempty,max=100"`
	LicenseNumber  string `json:"license_number" validate:"omitempty,max=50"`
	Phone          string `json:"phone" validate:"omitempty,max=15"`
	Email          string `json:"email" validate:"omitempty,email,max=100"`
	DepartmentID   *uint  `json:"department_id" validate:"omitempty"`
	IsActive       bool   `json:"is_active" validate:"omitempty"`
}

type ListDoctorQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id specialization created_at"`
}

func (q *ListDoctorQuery) SetDoctorDefaults() {
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

func (q *ListDoctorQuery) GetDoctorOffSet() int {
	return (q.Page - 1) * q.Limit
}
