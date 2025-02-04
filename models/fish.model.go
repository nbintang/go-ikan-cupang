package models

import "time"

type Fish struct {
	ID          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"unique;not null"`
	Description *string `gorm:"type:varchar(400)"`
	Price       float64 `gorm:"not null"`
	Stock       int     `gorm:"not null"`
	Image       *string
	CategoryID  int         `gorm:"not null"`
	Category    Category    `gorm:"foreignKey:CategoryID"`
	OrderItem   []OrderItem `gorm:"foreignKey:FishID"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
}
