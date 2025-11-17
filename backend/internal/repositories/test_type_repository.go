package repositories

import (
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type TestTypeRepository interface {
	Crete(testType *models.TestType) error
	Update(testType *models.TestType) error
	Delete(id uint) error
	HardDelete(id uint) error
	Restore(id uint) error

	FindById(id uint) (*models.TestType, error)
	FindAll(query *validators.ListTestTypeQuery) ([]models.TestType, int64, error)
	FindAllDelete(query *validators.ListTestTypeQuery) ([]models.TestType, int64, error)

	ExistsByCode(testTypeCode string) (bool, error)
}

type testTypeRepository struct {
	db *gorm.DB
}

func NewTestTypeRepository(db *gorm.DB) TestTypeRepository {
	return &testTypeRepository{
		db: db,
	}
}

func (r *testTypeRepository) Crete(testType *models.TestType) error {
	return r.db.Create(testType).Error
}

func (r *testTypeRepository) Update(testType *models.TestType) error {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) Delete(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) HardDelete(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) Restore(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) FindById(id uint) (*models.TestType, error) {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) FindAll(query *validators.ListTestTypeQuery) ([]models.TestType, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) FindAllDelete(query *validators.ListTestTypeQuery) ([]models.TestType, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *testTypeRepository) ExistsByCode(testTypeCode string) (bool, error) {
	var count int64
	err := r.db.Model(&models.TestType{}).Where("code = ?", testTypeCode).Count(&count).Error
	return count > 0, err
}
