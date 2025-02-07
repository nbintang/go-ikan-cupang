package models

import "time"

type Category struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"unique;type:varchar(80);not null"`
	Fishes    []Fish    `gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}