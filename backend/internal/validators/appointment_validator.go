package validators

type CreateAppointmentRequest struct {
	PatientID       uint   `json:"patient_id" validate:"required"`
	DoctorID        uint   `json:"doctor_id" validate:"required"`
	AppointmentDate string `json:"appointment_date" validate:"required"`
	AppointmentTime string `json:"appointment_time" validate:"required"` // HH:MM:SS
	DurationMinutes int    `json:"duration_minutes" validate:"required,min=1"`
	Reason          string `json:"reason" validate:"omitempty"`
	Status          string `json:"status" validate:"required,oneof=scheduled confirmed in_progress completed cancelled no_show"`
	Notes           string `json:"notes" validate:"omitempty"`
}

type UpdateAppointmentRequest struct {
	PatientID       uint   `json:"patient_id" validate:"omitempty"`
	DoctorID        uint   `json:"doctor_id" validate:"omitempty"`
	AppointmentDate string `json:"appointment_date" validate:"omitempty"`
	AppointmentTime string `json:"appointment_time" validate:"omitempty"` // HH:MM:SS
	DurationMinutes int    `json:"duration_minutes" validate:"omitempty,min=1"`
	Reason          string `json:"reason" validate:"omitempty"`
	Status          string `json:"status" validate:"omitempty,oneof=scheduled confirmed in_progress completed cancelled no_show"`
	Notes           string `json:"notes" validate:"omitempty"`
}

type ListAppointmentQuery struct {
	Page   int    `query:"page" validate:"omitempty,min=1"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Search string `query:"search" validate:"omitempty,max=100"`
	Sort   string `query:"sort" validate:"omitempty,oneof=asc desc"`
	SortBy string `query:"sort_by" validate:"omitempty,oneof=id appointment_date status created_at"`
}

func (q *ListAppointmentQuery) SetAppointmentDefaults() {
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

func (q *ListAppointmentQuery) GetAppointmentOffSet() int {
	return (q.Page - 1) * q.Limit
}
