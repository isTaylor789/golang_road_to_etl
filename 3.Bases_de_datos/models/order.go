package models

import "time"

type Order struct {
	ID        uint    `gorm:"primaryKey"`
	Product   string  `gorm:"type:varchar(100);not null"`
	Price     float64 `gorm:"not null"`
	UserID    uint    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
