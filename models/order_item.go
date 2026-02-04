package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrderID         uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	MenuItemID      uuid.UUID `gorm:"type:uuid;not null" json:"menu_item_id"`
	Quantity        int       `gorm:"not null" json:"quantity"`
	PriceAtPurchase float64   `gorm:"type:decimal(10,2);not null" json:"price_at_purchase"`

	MenuItem *MenuItem `gorm:"foreignKey:MenuItemID" json:"menu_item,omitempty"`
}