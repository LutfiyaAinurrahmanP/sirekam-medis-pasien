package repositories

import (
	"fmt"
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	Create(department *models.Department) error
	Update(department *models.Department) error
	Delete(id uint) error
	HardDelete(id uint) error
	Restore(id uint) error

	FindById(id uint) (*models.Department, error)
	FindAll(query *validators.ListDepartmentQuery) ([]models.Department, int64, error)
	FindAllDelete(query *validators.ListDepartmentQuery) ([]models.Department, int64, error)

	ExistsByCode(code string) (bool, error)
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{
		db: db,
	}
}

func (r *departmentRepository) Create(department *models.Department) error {
	return r.db.Create(department).Error
}

func (r *departmentRepository) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *departmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}

func (r *departmentRepository) HardDelete(id uint) error {
	return r.db.Unscoped().Delete(&models.Department{}, id).Error
}

func (r *departmentRepository) Restore(id uint) error {
	return r.db.Model(&models.Department{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *departmentRepository) FindById(id uint) (*models.Department, error) {
	var department models.Department
	err := r.db.First(&department, id).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}

func (r *departmentRepository) FindAll(query *validators.ListDepartmentQuery) ([]models.Department, int64, error) {
	var department []models.Department
	var total int64

	db := r.db.Model(&models.Department{})

	if query.Search != ""{
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(name) LIKE ? OR LOWER(code) LIKE ? OR LOWER(floor_location) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count departments: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetDepartmentOffSet())

	if err := db.Find(&department).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch department: %w", err)
	}
	return department, total, nil
}

func (r *departmentRepository) FindAllDelete(query *validators.ListDepartmentQuery) ([]models.Department, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *departmentRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Department{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}
