package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	PaymentStatusUnpaid    = "unpaid"
	PaymentStatusPartial   = "partial"
	PaymentStatusPaid      = "paid"
	PaymentStatusCancelled = "cancelled"
)

const (
	PaymentMethodCash       = "cash"
	PaymentMethodDebitCard  = "debit_card"
	PaymentMethodCreditCard = "credit_card"
	PaymentMethodInsurance  = "insurance"
	PaymentMethodTransfer   = "transfer"
	PaymentMethodOther      = "other"
)

type Billing struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	PatientID       uint           `gorm:"not null;index" json:"patient_id" validate:"required"`
	MedicalRecordID *uint          `gorm:"index" json:"medical_record_id,omitempty" validate:"omitempty"`
	InvoiceNumber   string         `gorm:"unique;not null;size:50;index" json:"invoice_number" validate:"required,max=50"`
	InvoiceDate     time.Time      `gorm:"not null;type:date;index" json:"invoice_date" validate:"required"`
	TotalAmount     float64        `gorm:"not null;type:decimal(12,2)" json:"total_amount" validate:"required,min=0"`
	PaidAmount      float64        `gorm:"not null;type:decimal(12,2);default:0" json:"paid_amount" validate:"omitempty"`
	DiscountAmount  *float64       `gorm:"type:decimal(12,2)" json:"discount_amount,omitempty" validate:"omitempty"`
	PaymentStatus   string         `gorm:"type:enum('unpaid','partial','paid','cancelled');not null;default:'unpaid';index" json:"payment_status" validate:"required,oneof=unpaid partial paid cancelled"`
	PaymentMethod   string         `gorm:"type:enum('cash','debit_card','credit_card','insurance','transfer','other')" json:"payment_method,omitempty" validate:"omitempty,oneof=cash debit_card credit_card insurance transfer other"`
	PaymentDate     *time.Time     `gorm:"type:datetime" json:"payment_date,omitempty" validate:"omitempty"`
	Notes           string         `gorm:"type:text" json:"notes,omitempty" validate:"omitempty"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Patient       *Patient       `gorm:"foreignKey:PatientID;references:ID" json:"patient,omitempty"`
	MedicalRecord *MedicalRecord `gorm:"foreignKey:MedicalRecordID;references:ID" json:"medical_record,omitempty"`
	Items         []BillingItem  `gorm:"foreignKey:BillingID;references:ID" json:"items,omitempty"`
}

func (Billing) TableName() string {
	return "billing"
}

func (b *Billing) BeforeCreate(tx *gorm.DB) error {
	if b.PaymentStatus == "" {
		b.PaymentStatus = PaymentStatusUnpaid
	}

	now := time.Now()
	if b.InvoiceDate.IsZero() {
		b.InvoiceDate = now
	}
	if b.CreatedAt.IsZero() {
		b.CreatedAt = now
	}
	if b.UpdatedAt.IsZero() {
		b.UpdatedAt = now
	}
	return nil
}

func (b *Billing) GetRemainingAmount() float64 {
	return b.TotalAmount - b.PaidAmount
}

func (b *Billing) GetNetAmount() float64 {
	discount := 0.0
	if b.DiscountAmount != nil {
		discount = *b.DiscountAmount
	}
	return b.TotalAmount - discount
}

func (b *Billing) AddPayment(amount float64) {
	b.PaidAmount += amount
	b.UpdatePaymentStatus()
}

func (b *Billing) UpdatePaymentStatus() {
	netAmount := b.GetNetAmount()

	if b.PaidAmount >= netAmount {
		b.PaymentStatus = PaymentStatusPaid
		now := time.Now()
		b.PaymentDate = &now
	} else if b.PaidAmount > 0 {
		b.PaymentStatus = PaymentStatusPartial
	} else {
		b.PaymentStatus = PaymentStatusUnpaid
	}
}

func (b *Billing) Cancel() {
	b.PaymentStatus = PaymentStatusCancelled
}

func (b *Billing) IsPaid() bool {
	return b.PaymentStatus == PaymentStatusPaid
}

func ValidatePaymentStatus(status string) bool {
	validStatuses := GetAvailablePaymentStatuses()
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}

func ValidatePaymentMethod(method string) bool {
	validMethods := GetAvailablePaymentMethods()
	for _, m := range validMethods {
		if m == method {
			return true
		}
	}
	return false
}

func GetAvailablePaymentStatuses() []string {
	return []string{
		PaymentStatusUnpaid,
		PaymentStatusPartial,
		PaymentStatusPaid,
		PaymentStatusCancelled,
	}
}

func GetAvailablePaymentMethods() []string {
	return []string{
		PaymentMethodCash,
		PaymentMethodDebitCard,
		PaymentMethodCreditCard,
		PaymentMethodInsurance,
		PaymentMethodTransfer,
		PaymentMethodOther,
	}
}
