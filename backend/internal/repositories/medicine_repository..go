package repositories

import (
	"fmt"
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Create(medicine *models.Medicine) error
	Update(medicine *models.Medicine) error
	Delete(id uint) error
	HardDelete(id uint) error
	Restore(id uint) error

	FindByID(id uint) (*models.Medicine, error)
	FindAll(query *validators.ListMedicineQuery) ([]models.Medicine, int64, error)
	FindAllDelete(query *validators.ListMedicineQuery) ([]models.Medicine, int64, error)
}

type medicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) MedicineRepository {
	return &medicineRepository{
		db: db,
	}
}

func (r *medicineRepository) Create(medicine *models.Medicine) error {
	return r.db.Create(medicine).Error
}

func (r *medicineRepository) Update(medicine *models.Medicine) error {
	return r.db.Save(medicine).Error
}

func (r *medicineRepository) Delete(id uint) error {
	return r.db.Delete(&models.Medicine{}, id).Error
}

func (r *medicineRepository) HardDelete(id uint) error {
	return r.db.Unscoped().Delete(&models.Medicine{}, id).Error
}

func (r *medicineRepository) Restore(id uint) error {
	return r.db.Model(&models.Medicine{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *medicineRepository) FindByID(id uint) (*models.Medicine, error) {
	var medicine models.Medicine
	err := r.db.First(&medicine, id).Error
	if err != nil {
		return nil, err
	}
	return &medicine, nil
}

func (r *medicineRepository) FindAll(query *validators.ListMedicineQuery) ([]models.Medicine, int64, error) {
	var medicines []models.Medicine
	var total int64

	db := r.db.Model(&models.Medicine{})
	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(name) LIKE ? OR LOWER(generic_name) LIKE ? OR LOWER(brand_name) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count medicine: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetMedicineOffSet())

	if err := db.Find(&medicines).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch medicine: %w", err)
	}

	return medicines, total, nil
}

func (r *medicineRepository) FindAllDelete(query *validators.ListMedicineQuery) ([]models.Medicine, int64, error) {
	panic("not implemented") // TODO: Implement
}
