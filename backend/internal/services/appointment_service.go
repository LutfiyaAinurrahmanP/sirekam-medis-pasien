package services

import (
	"fmt"
	"time"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
)

type AppointmentService interface {
	CreateAppointment(req *validators.CreateAppointmentRequest) (*models.Appointment, error)
	UpdateAppointment(id uint, req *validators.CreateAppointmentRequest) (*models.Appointment, error)
	DeleteAppointment(id uint) error
	HardDeleteAppointment(id uint) error
	RestoreAppointment(id uint) error

	GetAppointmentByID(id uint) (*models.Appointment, error)
	GetAllAppointment(query *validators.ListAppointmentQuery) ([]models.Appointment, *utils.PaginationMeta, error)
	GetAllDeleteAppointment(query *validators.ListAppointmentQuery) ([]models.Appointment, *utils.PaginationMeta, error)
}

type appointmentService struct {
	appointmentRepo repositories.AppointmentRepository
}

func NewAppointmentService(appointmentRepo repositories.AppointmentRepository) AppointmentService {
	return &appointmentService{
		appointmentRepo: appointmentRepo,
	}
}

func (s *appointmentService) CreateAppointment(req *validators.CreateAppointmentRequest) (*models.Appointment, error) {
	appointmentDate, err := time.Parse("2006-01-02", req.AppointmentDate)
	if err != nil {
		return nil, fmt.Errorf("invalid appointment_date format")
	}

	appointment := &models.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: appointmentDate,
		AppointmentTime: req.AppointmentTime,
		DurationMinutes: req.DurationMinutes,
		Reason:          req.Reason,
		Status:          req.Status,
		Notes:           req.Notes,
	}

	if err := s.appointmentRepo.Create(appointment); err != nil {
		return nil, fmt.Errorf("failed to created appointment: %w", err)
	}

	return appointment, nil
	// return nil, nil
}

func (s *appointmentService) UpdateAppointment(id uint, req *validators.CreateAppointmentRequest) (*models.Appointment, error) {
	panic("not implemented") // TODO: Implement
}

func (s *appointmentService) DeleteAppointment(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *appointmentService) HardDeleteAppointment(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *appointmentService) RestoreAppointment(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *appointmentService) GetAppointmentByID(id uint) (*models.Appointment, error) {
	panic("not implemented") // TODO: Implement
}

func (s *appointmentService) GetAllAppointment(query *validators.ListAppointmentQuery) ([]models.Appointment, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}

func (s *appointmentService) GetAllDeleteAppointment(query *validators.ListAppointmentQuery) ([]models.Appointment, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}
