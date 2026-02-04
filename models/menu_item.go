package models

import (
	"time"

	"github.com/google/uuid"
)

type MenuItem struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CanteenID uuid.UUID `gorm:"type:uuid;not null" json:"canteen_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock     int       `gorm:"not null;default:0" json:"stock"`
	Category  string    `gorm:"type:varchar(100);not null" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Canteen *Canteen `gorm:"foreignKey:CanteenID" json:"-"`
}