package models

import "time"

type VerificationToken struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Token     string    `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	UserID    int       `gorm:"not null;index"` // Tambahkan index agar lebih optimal
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
