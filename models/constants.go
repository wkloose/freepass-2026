package models

type UserRole string

const (
	RoleAdmin   UserRole = "ADMIN"
	RoleOwner   UserRole = "CANTEEN"
	RoleCustomer UserRole = "CUSTOMER"
)

type PaymentStatus string

const (
	PaymentStatusUnpaid PaymentStatus = "UNPAID"
	PaymentStatusPaid   PaymentStatus = "PAID"
)

type OrderStatus string

const (
	StatusUnpaid    OrderStatus = "UNPAID"
	StatusPaid      OrderStatus = "PAID"
	StatusWaiting   OrderStatus = "WAITING"
	StatusCooking   OrderStatus = "COOKING"
	StatusReady     OrderStatus = "READY"
	StatusCompleted OrderStatus = "COMPLETED"
	StatusCancelled OrderStatus = "CANCELLED"
)