package repositories

import (
	"fmt"
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	Update(user *models.User) error
	Delete(id uint) error
	HardDelete(id uint) error
	Restore(id uint) error

	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	FindAll(query *validators.ListUserQuery) ([]models.User, int64, error)
	FindAllDelete(query *validators.ListUserQuery) ([]models.User, int64, error)

	ExistsByUsername(username string) (bool, error)
	ExistsByEmail(email string) (bool, error)
	ExistsByPhone(phone string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) HardDelete(id uint) error {
	return r.db.Unscoped().Delete(&models.User{}, id).Error
}

func (r *userRepository) Restore(id uint) error {
	return r.db.Model(&models.User{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindAll(query *validators.ListUserQuery) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	db := r.db.Model(&models.User{})

	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(username) LIKE ? OR LOWER(email) LIKE ? OR LOWER(phone) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if query.Role != "" {
		db = db.Where("role = ?", query.Role)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetOffSet())

	if err := db.Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch user: %w", err)
	}
	return users, total, nil
}

func (r *userRepository) FindAllDelete(query *validators.ListUserQuery) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	db := r.db.Unscoped().Model(&models.User{}).Where("deleted_at IS NOT NULL")

	if query.Search != "" {
		searchPattern := "%" + strings.ToLower(query.Search) + "%"
		db = db.Where(
			"LOWER(username) LIKE ? OR LOWER(email) LIKE ? OR LOWER(PHONE) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if query.Role != "" {
		db = db.Where("role = ?", query.Role)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count deleted users: %w", err)
	}

	orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	db = db.Order(orderClause)

	db = db.Limit(query.Limit).Offset(query.GetOffSet())

	if err := db.Find(&users).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to fetch deleted users: %w", err)
	}

	return users, total, nil
}

func (r *userRepository) ExistsByUsername(username string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("username = ?", username).Count(&count).Error
	return count > 0, err
}

func (r *userRepository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *userRepository) ExistsByPhone(phone string) (bool, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("phone = ?", phone).Count(&count).Error
	return count > 0, err
}
