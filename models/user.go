package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"`
}

type Session struct {
	ID        string `gorm:"primaryKey"`
	UserID    uint
	User      User
	IPAddress string
	UserAgent string
	ExpiresAt time.Time
	CreatedAt time.Time
}
