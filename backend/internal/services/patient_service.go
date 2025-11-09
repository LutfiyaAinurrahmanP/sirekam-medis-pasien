package services

import (
	"errors"
	"fmt"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/repositories"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/utils"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type PatientService interface {
	CreatePatient(req *validators.CreatePatientRequest) (*models.Patient, error)
	UpdatePatient(id uint, req *validators.UpdatePatientRequest) (*models.Patient, error)
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

func (s *patientService) UpdatePatient(id uint, req *validators.UpdatePatientRequest) (*models.Patient, error) {
	patient, err := s.patientRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("patient not found")
		}
		return nil, fmt.Errorf("failed to find patient: %w", err)
	}

	if req.PatientCode != "" && req.PatientCode != patient.PatientCode {
		exists, err := s.patientRepo.ExistsByPatientCode(req.PatientCode)
		if err != nil {
			return nil, fmt.Errorf("failed to check patient code: %w", err)
		}
		if exists {
			return nil, errors.New("patient code already exists")
		}
		patient.PatientCode = req.PatientCode
	}

	if req.UserID != nil && req.UserID != patient.UserID {
		patient.UserID = req.UserID
	}

	if req.FullName != "" && req.FullName != patient.FullName {
		patient.FullName = req.FullName
	}

	if !req.DateOfBirth.IsZero() && req.DateOfBirth != patient.DateOfBirth {
		patient.DateOfBirth = req.DateOfBirth
	}

	if req.Gender != "" && req.Gender != patient.Gender {
		patient.Gender = req.Gender
	}

	if req.BloodType != "" && req.BloodType != patient.BloodType {
		patient.BloodType = req.BloodType
	}

	if req.Phone != "" && req.Phone != patient.Phone {
		patient.Phone = req.Phone
	}

	if req.Email != "" && req.Email != patient.Email {
		patient.Email = req.Email
	}

	if req.Address != "" && req.Address != patient.Address {
		patient.Address = req.Address
	}

	if req.EmergencyContactName != "" && req.EmergencyContactName != patient.EmergencyContactName {
		patient.EmergencyContactName = req.EmergencyContactName
	}

	if req.EmergencyContactPhone != "" && req.EmergencyContactPhone != patient.EmergencyContactPhone {
		patient.EmergencyContactPhone = req.EmergencyContactPhone
	}

	if req.InsuranceNumber != "" && req.InsuranceNumber != patient.InsuranceNumber {
		patient.InsuranceNumber = req.InsuranceNumber
	}

	if req.InsuranceProvider != "" && req.InsuranceProvider != patient.InsuranceProvider {
		patient.InsuranceProvider = req.InsuranceProvider
	}

	if req.Allergies != "" && req.Allergies != patient.Allergies {
		patient.Allergies = req.Allergies
	}

	if err := s.patientRepo.Update(patient); err != nil {
		return nil, fmt.Errorf("failed to updated patient: %w", err)
	}
	return patient, err
}

func (s *patientService) DeletePatient(id uint) error {
	_, err := s.patientRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return errors.New("patient not found")
		}
		return fmt.Errorf("failed to find patient: %w", err)
	}

	if err := s.patientRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete patient: %w", err)
	}
	return nil
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
