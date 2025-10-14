package validators

type CreateRoomRequest struct {
	RoomNumber    string   `json:"room_number" validate:"required,max=20"`
	RoomType      string   `json:"room_type" validate:"required,oneof=vip class_1 class_2 class_3 icu emergency"`
	DepartmentID  *uint    `json:"department_id" validate:"omitempty"`
	BedCapacity   int      `json:"bed_capacity" validate:"required,min=1"`
	AvailableBeds int      `json:"available_beds" validate:"required"`
	PricePerDay   *float64 `json:"price_per_day" validate:"omitempty"`
	IsActive      bool     `json:"is_active" validate:"required"`
}

type UpdateRoomRequest struct {
	RoomNumber    string   `json:"room_number" validate:"omitempty,max=20"`
	RoomType      string   `json:"room_type" validate:"omitempty,oneof=vip class_1 class_2 class_3 icu emergency"`
	DepartmentID  *uint    `json:"department_id" validate:"omitempty"`
	BedCapacity   int      `json:"bed_capacity" validate:"omitempty,min=1"`
	AvailableBeds int      `json:"available_beds" validate:"omitempty"`
	PricePerDay   *float64 `json:"price_per_day" validate:"omitempty"`
	IsActive      bool     `json:"is_active" validate:"omitempty"`
}

type ListRoomQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id room_type available_beds is_active created_at"`
}

func (q *ListRoomQuery) SetRoomDefaults() {
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

func (q *ListRoomQuery) GetRoomOffSet() int {
	return (q.Page - 1) * q.Limit
}
