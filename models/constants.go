package models

type UserRole string

const (
	RoleAdmin   UserRole = "ADMIN"
	RoleOwner   UserRole = "OWNER"
	RoleCustomer UserRole = "CUSTOMER"
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