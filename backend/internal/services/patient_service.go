package services

import (
	"errors"
	"fmt"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
)

type PatientService interface {
	CreatePatient(req *validators.CreatePatientRequest) (*models.Patient, error)
	UpdatePatient(id, req *validators.UpdatePatientRequest) (*models.Patient, error)
	DeletePatient(id uint) error
	HardDeletePatient(id uint) error
	Restore(id uint) error

	GetPatientByID(id uint) (*models.Patient, error)
	GetAllPatient(query *validators.ListPatientQuery) ([]models.Patient, *utils.PaginationMeta, error)
	GetAllDeletePatient(query *validators.ListPatientQuery) ([]models.Patient, *utils.PaginationMeta, error)
}

type patientService struct {
	patientRepo repositories.PatientRepository
}

func NewPatientService(patientRepo repositories.PatientRepository) PatientService {
	return &patientService{
		patientRepo: patientRepo,
	}
}

func (s *patientService) CreatePatient(req *validators.CreatePatientRequest) (*models.Patient, error) {
	exists, err := s.patientRepo.ExistsByPatientCode(req.PatientCode)
	if err != nil {
		return nil, fmt.Errorf("failed to check patient code: %w", err)
	}

	if exists {
		return nil, errors.New("patient code already exists")
	}

	patient := &models.Patient{
		UserID:                req.UserID,
		PatientCode:           req.PatientCode,
		FullName:              req.FullName,
		DateOfBirth:           req.DateOfBirth,
		Gender:                req.Gender,
		BloodType:             req.BloodType,
		Phone:                 req.Phone,
		Email:                 req.Email,
		Address:               req.Address,
		EmergencyContactName:  req.EmergencyContactName,
		EmergencyContactPhone: req.EmergencyContactPhone,
		InsuranceNumber:       req.InsuranceNumber,
		InsuranceProvider:     req.InsuranceProvider,
		Allergies:             req.Allergies,
	}

	if err := s.patientRepo.Create(patient); err != nil {
		return nil, fmt.Errorf("failed to created patient: %w", err)
	}

	return patient, nil
}

func (s *patientService) UpdatePatient(id *validators.UpdatePatientRequest, req *validators.UpdatePatientRequest) (*models.Patient, error) {
	panic("not implemented") // TODO: Implement
}

func (s *patientService) DeletePatient(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *patientService) HardDeletePatient(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *patientService) Restore(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (s *patientService) GetPatientByID(id uint) (*models.Patient, error) {
	panic("not implemented") // TODO: Implement
}

func (s *patientService) GetAllPatient(query *validators.ListPatientQuery) ([]models.Patient, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}

func (s *patientService) GetAllDeletePatient(query *validators.ListPatientQuery) ([]models.Patient, *utils.PaginationMeta, error) {
	panic("not implemented") // TODO: Implement
}
