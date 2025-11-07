package services

import (
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
        panic("not implemented") // TODO: Implement
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