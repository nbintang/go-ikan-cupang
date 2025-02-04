package models

import (
	"time"
)

type Payment struct {
    ID        uint          `gorm:"primaryKey"`
    OrderID   uint          `gorm:"not null"`
    Amount    float64       `gorm:"not null"`
    Status    PaymentStatus `gorm:"type:ENUM('PENDING', 'PAID', 'FAILED');default:'PENDING'"`
    CreatedAt time.Time
    UpdatedAt time.Time
}