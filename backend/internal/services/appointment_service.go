package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type AppointmentService interface {
	CreateAppointment(req *validators.CreateAppointmentRequest) (*models.Appointment, error)
	UpdateAppointment(id uint, req *validators.UpdateAppointmentRequest) (*models.Appointment, error)
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
}

func (s *appointmentService) UpdateAppointment(id uint, req *validators.UpdateAppointmentRequest) (*models.Appointment, error) {
	appointment, err := s.appointmentRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("appointment not found")
		}
		return nil, fmt.Errorf("failed to find appointment: %w", err)
	}
	if err := s.appointmentRepo.Update(appointment); err != nil {
		return nil, fmt.Errorf("failed to updated appointment: %w", err)
	}
	return appointment, nil
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
