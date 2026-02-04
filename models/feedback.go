package models

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null;unique" json:"order_id"`
	Rating    int       `gorm:"not null" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `json:"created_at"`

	Order *Order `gorm:"foreignKey:OrderID" json:"-"`
}