package repositories

import (
	"fmt"
	"strings"

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
	return r.db.Model(&models.Patient{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
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
	var patient []models.Patient
	var total int64

	db := r.db.Model(&models.Patient{})

	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search)+ "%"
		db = db.Where(
			"LOWER(patient_code) LIKE ? OR LOWER(full_name) LIKE ? OR LOWER(phone) LIKE ? OR LOWER(email) LIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern, 
		)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count patient: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetPatientOffSet())

	if err := db.Find(&patient).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch patient: %w", err)
	}

	return patient, total, nil
}

func (r *patientRepository) FindAllDelete(query *validators.ListPatientQuery) ([]models.Patient, int64, error) {
	var patient []models.Patient
	var total int64

	db := r.db.Unscoped().Model(&models.Patient{}).Where("deleted_at IS NOT NULL")

	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(patient_code) LIKE ? OR LOWER(full_name) LIKE ? OR LOWER(phone) LIKE ? OR LOWER(email) LIKE ?",
			searchPattern, searchPattern, searchPattern, searchPattern, 
		)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count patients: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetPatientOffSet())

	if err := db.Find(&patient).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch patients: %w", err)
	}
	return patient, total, nil
}

func (r *patientRepository) ExistsByPatientCode(patientCode string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Patient{}).Where("patient_code = ?", patientCode).Count(&count).Error
	return count > 0, err
}
