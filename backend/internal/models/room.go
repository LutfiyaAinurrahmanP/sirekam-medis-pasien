package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	RoomTypeVIP       = "vip"
	RoomTypeClass1    = "class_1"
	RoomTypeClass2    = "class_2"
	RoomTypeClass3    = "class_3"
	RoomTypeICU       = "icu"
	RoomTypeEmergency = "emergency"
)

type Room struct {
	ID            uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	RoomNumber    string         `gorm:"unique;not null;size:20;index" json:"room_number" validate:"required,max=20"`
	RoomType      string         `gorm:"type:enum('vip','class_1','class_2','class_3','icu','emergency');not null;index" json:"room_type" validate:"required,oneof=vip class_1 class_2 class_3 icu emergency"`
	DepartmentID  *uint          `gorm:"index" json:"department_id,omitempty" validate:"omitempty"`
	BedCapacity   int            `gorm:"not null" json:"bed_capacity" validate:"required,min=1"`
	AvailableBeds int            `gorm:"not null" json:"available_beds" validate:"required"`
	PricePerDay   *float64       `gorm:"type:decimal(10,2)" json:"price_per_day,omitempty" validate:"omitempty"`
	IsActive      bool           `gorm:"not null;default:true;index" json:"is_active" validate:"required"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Department       *Department       `gorm:"foreignKey:DepartmentID;references:ID" json:"department,omitempty"`
	Hospitalizations []Hospitalization `gorm:"foreignKey:RoomID;references:ID" json:"hospitalizations,omitempty"`
}

func (Room) TableName() string {
	return "rooms"
}

func (r *Room) BeforeCreate(tx *gorm.DB) error {
	if r.IsActive {
		r.IsActive = true
	}

	// Set available beds to bed capacity if not set
	if r.AvailableBeds == 0 {
		r.AvailableBeds = r.BedCapacity
	}

	now := time.Now()
	if r.CreatedAt.IsZero() {
		r.CreatedAt = now
	}
	if r.UpdatedAt.IsZero() {
		r.UpdatedAt = now
	}
	return nil
}

func (r *Room) IsAvailable() bool {
	return r.AvailableBeds > 0 && r.IsActive
}

func (r *Room) IsFull() bool {
	return r.AvailableBeds <= 0
}

func (r *Room) OccupyBed() error {
	if r.AvailableBeds <= 0 {
		return gorm.ErrInvalidData
	}
	r.AvailableBeds--
	return nil
}

func (r *Room) ReleaseBed() error {
	if r.AvailableBeds >= r.BedCapacity {
		return gorm.ErrInvalidData
	}
	r.AvailableBeds++
	return nil
}

func (r *Room) GetOccupancyRate() float64 {
	if r.BedCapacity == 0 {
		return 0
	}
	occupied := r.BedCapacity - r.AvailableBeds
	return (float64(occupied) / float64(r.BedCapacity)) * 100
}

func ValidateRoomType(roomType string) bool {
	validTypes := GetAvailableRoomTypes()
	for _, t := range validTypes {
		if t == roomType {
			return true
		}
	}
	return false
}

func GetAvailableRoomTypes() []string {
	return []string{
		RoomTypeVIP,
		RoomTypeClass1,
		RoomTypeClass2,
		RoomTypeClass3,
		RoomTypeICU,
		RoomTypeEmergency,
	}
}
