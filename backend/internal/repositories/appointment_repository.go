package repositories

import (
	"fmt"
	"strings"

	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/models"
	"github.com/LutfiyaAinurrahmanP/sirekam-medis-pasien/internal/validators"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Create(appointment *models.Appointment) error
	Update(apppointment *models.Appointment) error
	Delete(id uint) error
	HardDelete(id uint) error
	Restore(id uint) error

	FindById(id uint) (*models.Appointment, error)
	FindAll(query *validators.ListAppointmentQuery) ([]models.Appointment, int64, error)
	FindAllDelete(query *validators.ListAppointmentQuery) ([]models.Appointment, int64, error)
}

type appointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &appointmentRepository{
		db: db,
	}
}

func (r *appointmentRepository) Create(appointment *models.Appointment) error {
        return r.db.Create(appointment).Error
}

func (r *appointmentRepository) Update(apppointment *models.Appointment) error {
        return r.db.Save(apppointment).Error
}

func (r *appointmentRepository) Delete(id uint) error {
        return r.db.Delete(&models.Appointment{}, id).Error
}

func (r *appointmentRepository) HardDelete(id uint) error {
        return r.db.Unscoped().Delete(&models.Appointment{}, id).Error
}

func (r *appointmentRepository) Restore(id uint) error {
        return r.db.Model(&models.Appointment{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil).Error
}

func (r *appointmentRepository) FindById(id uint) (*models.Appointment, error) {
        var appointment models.Appointment
		err := r.db.First(&appointment, id).Error
		if err != nil {
			return nil, err
		}
		return &appointment, nil
}

func (r *appointmentRepository) FindAll(query *validators.ListAppointmentQuery) ([]models.Appointment, int64, error) {
        var appointments []models.Appointment
		var total int64

		db := r.db.Model(&models.Appointment{})

		if query.Search != "" {
			searchPattern := "%" + strings.ToLower(query.Search) + "%"
			db = db.Where(
				"LOWER(appointment_date) LIKE ? OR LOWER(status) LIKE ?",
				searchPattern, searchPattern,
			)
		}

		if err := db.Count(&total).Error; err != nil {
			return nil, 0, fmt.Errorf("failed to count appointment: %w", err)
		}

		orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
		db = db.Order(orderClause)

		db = db.Limit(query.Limit).Offset(query.GetAppointmentOffSet())

		if err := db.Find(&appointments).Error; err != nil {
			return nil, 0, fmt.Errorf("failed to fetch appointment: %w", err)
		}

		return appointments, total, nil
}

func (r *appointmentRepository) FindAllDelete(query *validators.ListAppointmentQuery) ([]models.Appointment, int64, error) {
        var appointments []models.Appointment
		var total int64

		db := r.db.Unscoped().Model(&models.Appointment{}).Where("deleted_at IS NOT NULL")

		if query.Search != "" {
			searchPattern := "%" + strings.ToLower(query.Search) + "%"
			db = db.Where(
				"LOWER(appointment_date) LIKE ? OR LOWER(status) LIKE ?",
				searchPattern, searchPattern,
			)
		}

		if err := db.Count(&total).Error; err != nil {
			return nil, 0, fmt.Errorf("failed to count appointment: %w", err)
		}

		orderClause := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
		db = db.Order(orderClause)

		db = db.Limit(query.Limit).Offset(query.GetAppointmentOffSet())

		if err := db.Find(&appointments).Error; err != nil {
			return nil, 0, fmt.Errorf("failed to fetch appointment: %w", err)
		}

		return appointments, total, nil
}