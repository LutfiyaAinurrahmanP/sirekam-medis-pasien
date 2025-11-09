package repositories

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	Update(patient *models.Patient) error
	Delete(id uint) error
	HardDelete(id uint) error
	Restore(id uint) error

	FindById(id uint) (*models.Patient, error)
	FindAll(query *validators.ListPatientQuery) ([]models.Patient, int64, error)
	FindAllDelete(query *validators.ListPatientQuery) ([]models.Patient, int64, error)

	ExistsByPatientCode(patientCode string) (bool, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{
		db: db,
	}
}

func (r *patientRepository) Create(patient *models.Patient) error {
	return r.db.Create(patient).Error
}

func (r *patientRepository) Update(patient *models.Patient) error {
	return r.db.Save(patient).Error
}

func (r *patientRepository) Delete(id uint) error {
	return r.db.Delete(&models.Patient{}, id).Error
}

func (r *patientRepository) HardDelete(id uint) error {
	return r.db.Unscoped().Delete(&models.Patient{}, id).Error
}

func (r *patientRepository) Restore(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (r *patientRepository) FindById(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.db.First(&patient, id).Error
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func (r *patientRepository) FindAll(query *validators.ListPatientQuery) ([]models.Patient, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *patientRepository) FindAllDelete(query *validators.ListPatientQuery) ([]models.Patient, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *patientRepository) ExistsByPatientCode(patientCode string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Patient{}).Where("patient_code = ?", patientCode).Count(&count).Error
	return count > 0, err
}
