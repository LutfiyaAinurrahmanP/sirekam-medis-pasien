package repositories

import (
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
	panic("not implemented") // TODO: Implement
}

func (r *medicineRepository) Restore(id uint) error {
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
}

func (r *medicineRepository) FindAllDelete(query *validators.ListMedicineQuery) ([]models.Medicine, int64, error) {
	panic("not implemented") // TODO: Implement
}
