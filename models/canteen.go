package models

import (
	"time"

	"github.com/google/uuid"
)

type Canteen struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OwnerID     uuid.UUID `gorm:"type:uuid;unique;not null" json:"owner_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	IsOpen      bool      `gorm:"default:false" json:"is_open"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Owner     *User      `gorm:"foreignKey:OwnerID" json:"owner,omitempty"`
	MenuItems []MenuItem `gorm:"foreignKey:CanteenID;constraint:OnDelete:CASCADE;" json:"menu_items,omitempty"`
	Orders    []Order    `gorm:"foreignKey:CanteenID" json:"orders,omitempty"`
}