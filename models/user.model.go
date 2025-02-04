package models

import (
	"time"

	_"gorm.io/gorm"
)

type User struct {
	ID         int                 `gorm:"primaryKey;autoIncrement"`
	Name       string              `gorm:"type:varchar(100);not null"`
	Email      string              `gorm:"unique;not null"`
	Role       Role                `gorm:"type:varchar(10);default:'USER'"` // Gunakan varchar(10) daripada role
	Orders     []Order             `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time           `gorm:"autoCreateTime"`
	UpdatedAt  time.Time           `gorm:"autoUpdateTime"`
	IsVerified bool                `gorm:"default:false"`
	Token      []VerificationToken `gorm:"foreignKey:UserID"`
}
