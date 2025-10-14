package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	BillingItemTypeConsultation = "consultation"
	BillingItemTypeMedicine     = "medicine"
	BillingItemTypeLabTest      = "lab_test"
	BillingItemTypeProcedure    = "procedure"
	BillingItemTypeRoom         = "room"
	BillingItemTypeOther        = "other"
)

type BillingItem struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	BillingID       uint           `gorm:"not null;index" json:"billing_id" validate:"required"`
	ItemType        string         `gorm:"type:enum('consultation','medicine','lab_test','procedure','room','other');not null;index" json:"item_type" validate:"required,oneof=consultation medicine lab_test procedure room other"`
	ItemDescription string         `gorm:"not null;size:255" json:"item_description" validate:"required,max=255"`
	Quantity        int            `gorm:"not null" json:"quantity" validate:"required,min=1"`
	UnitPrice       float64        `gorm:"not null;type:decimal(10,2)" json:"unit_price" validate:"required,min=0"`
	TotalPrice      float64        `gorm:"not null;type:decimal(10,2)" json:"total_price" validate:"omitempty,min=0"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Billing *Billing `gorm:"foreignKey:BillingID;references:ID" json:"billing,omitempty"`
}

func (BillingItem) TableName() string {
	return "billing_items"
}

func (b *BillingItem) BeforeCreate(tx *gorm.DB) error {
	// Auto calculate total price
	b.TotalPrice = float64(b.Quantity) * b.UnitPrice

	now := time.Now()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = now
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = now
	}
	return nil
}

func (b *BillingItem) BeforeSave(tx *gorm.DB) error {
	// Recalculate total price before saving
	b.TotalPrice = float64(b.Quantity) * b.UnitPrice
	return nil
}

func (b *BillingItem) CalculateTotalPrice() float64 {
	return float64(b.Quantity) * b.UnitPrice
}

func ValidateBillingItemType(itemType string) bool {
	validTypes := GetAvailableBillingItemTypes()
	for _, t := range validTypes {
		if t == itemType {
			return true
		}
	}
	return false
}

func GetAvailableBillingItemTypes() []string {
	return []string{
		BillingItemTypeConsultation,
		BillingItemTypeMedicine,
		BillingItemTypeLabTest,
		BillingItemTypeProcedure,
		BillingItemTypeRoom,
		BillingItemTypeOther,
	}
}
