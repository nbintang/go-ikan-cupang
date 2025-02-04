package models

import (
	"time"
)

type Order struct {
	ID         int         `gorm:"primaryKey;autoIncrement"`
	UserID     int         `gorm:"not null"`
	User       User        `gorm:"foreignKey:UserID"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	TotalPrice float64     `gorm:"not null"`
	Payments   []Payment   `gorm:"foreignKey:OrderID"`
	Status     OrderStatus `gorm:"type:varchar(20);default:'PENDING'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
