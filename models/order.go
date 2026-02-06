package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	UserID        uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
	CanteenID     uuid.UUID     `gorm:"type:uuid;not null" json:"canteen_id"`
	TotalPrice    float64       `gorm:"type:decimal(10,2);not null" json:"total_price"`
	
	PaymentStatus PaymentStatus `gorm:"type:varchar(20);default:'UNPAID'" json:"payment_status"`
	Status        OrderStatus   `gorm:"type:varchar(20);default:'WAITING'" json:"status"`
	
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`

	User       *User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Canteen    *Canteen    `gorm:"foreignKey:CanteenID" json:"canteen,omitempty"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"order_items,omitempty"`
	Feedback   *Feedback   `gorm:"foreignKey:OrderID" json:"feedback,omitempty"`
}