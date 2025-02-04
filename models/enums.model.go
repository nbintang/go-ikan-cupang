package models

type Role string

const (
	USER  Role = "USER"
	ADMIN Role = "ADMIN"
)

type OrderStatus string

const (
	PENDING   OrderStatus = "PENDING"
	PAID      OrderStatus = "PAID"
	SHIPPED   OrderStatus = "SHIPPED"
	COMPLETED OrderStatus = "COMPLETED"
	CANCELLED OrderStatus = "CANCELLED"
)

type PaymentStatus string

const (
	PAYMENT_PENDING PaymentStatus = "PENDING"
	SUCCESS         PaymentStatus = "SUCCESS"
	FAILED          PaymentStatus = "FAILED"
)
