package repositories

import (
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
	panic("not implemented") // TODO: Implement
}

func (r *departmentRepository) HardDelete(id uint) error {
	panic("not implemented") // TODO: Implement
}

func (r *departmentRepository) Restore(id uint) error {
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
}

func (r *departmentRepository) FindAllDelete(query *validators.ListDepartmentQuery) ([]models.Department, int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *departmentRepository) ExistsByCode(code string) (bool, error) {
	var count int64
	err := r.db.Model(&models.Department{}).Where("code = ?", code).Count(&count).Error
	return count > 0, err
}
