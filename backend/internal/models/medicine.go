package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	MedicineTypeTablet    = "tablet"
	MedicineTypeCapsule   = "capsule"
	MedicineTypeSyrup     = "syrup"
	MedicineTypeInjection = "injection"
	MedicineTypeOintment  = "ointment"
	MedicineTypeOther     = "other"
)

type Medicine struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"not null;size:200;index" json:"name" validate:"required"`
	GenericName   string         `gorm:"size:200" json:"generic_name,omitempty"`
	BrandName     string         `gorm:"size:200" json:"brand_name,omitempty"`
	Type          string         `gorm:"type:enum('tablet','capsule','syrup','injection','ointment','other');not null" json:"type" validate:"required"`
	Strength      string         `gorm:"size:50" json:"strength,omitempty"` // e.g., "500mg"
	Manufacturer  string         `gorm:"size:100" json:"manufacturer,omitempty"`
	Unit          string         `gorm:"size:20" json:"unit,omitempty"` // tablet, ml, mg
	StockQuantity int            `gorm:"not null;default:0" json:"stock_quantity"`
	Price         *float64       `gorm:"type:decimal(10,2)" json:"price,omitempty"`
	IsActive      bool           `gorm:"not null;default:true;index" json:"is_active"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	PrescriptionItems []PrescriptionItem `gorm:"foreignKey:MedicineID;references:ID" json:"prescription_items,omitempty"`
}

func (Medicine) TableName() string {
	return "medicines"
}

func (m *Medicine) BeforeCreate(tx *gorm.DB) error {
	if m.IsActive {
		m.IsActive = true
	}

	now := time.Now()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = now
	}
	return nil
}

func (m *Medicine) AddStock(quantity int) {
	m.StockQuantity += quantity
}

func (m *Medicine) ReduceStock(quantity int) error {
	if m.StockQuantity < quantity {
		return gorm.ErrInvalidData
	}
	m.StockQuantity -= quantity
	return nil
}

func (m *Medicine) IsOutOfStock() bool {
	return m.StockQuantity <= 0
}

func (m *Medicine) IsLowStock(threshold int) bool {
	return m.StockQuantity <= threshold
}

func ValidateMedicineType(medicineType string) bool {
	validTypes := GetAvailableMedicineTypes()
	for _, t := range validTypes {
		if t == medicineType {
			return true
		}
	}
	return false
}

func GetAvailableMedicineTypes() []string {
	return []string{
		MedicineTypeTablet,
		MedicineTypeCapsule,
		MedicineTypeSyrup,
		MedicineTypeInjection,
		MedicineTypeOintment,
		MedicineTypeOther,
	}
}