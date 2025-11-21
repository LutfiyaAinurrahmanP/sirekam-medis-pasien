package repositories

import (
	"fmt"
	"strings"

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
	var testType models.TestType
	err := r.db.First(&testType, id).Error
	if err != nil {
		return nil, err
	}
	return &testType, nil
}

func (r *testTypeRepository) FindAll(query *validators.ListTestTypeQuery) ([]models.TestType, int64, error) {
	var testType []models.TestType
	var total int64

	db := r.db.Model(&models.TestType{})

	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(name) LIKE ? OR LOWER(code) LIKE ? OR LOWER(category) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count test type: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetTestTypeOffSet())

	if err := db.Find(&testType).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch test type: %w", err)
	}

	return testType, total, nil
}

func (r *testTypeRepository) FindAllDelete(query *validators.ListTestTypeQuery) ([]models.TestType, int64, error) {
	var testType []models.TestType
	var total int64

	db := r.db.Unscoped().Model(&models.TestType{}).Where("deleted_at IS NOT NULL")

	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(name) LIKE ? OR LOWER(code) LIKE ? OR LOWER(category) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count test types: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetTestTypeOffSet())

	if err := db.Find(&testType).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch test types: %w", err)
	}

	return testType, total, nil
}

func (r *testTypeRepository) ExistsByCode(testTypeCode string) (bool, error) {
	var count int64
	err := r.db.Model(&models.TestType{}).Where("code = ?", testTypeCode).Count(&count).Error
	return count > 0, err
}
