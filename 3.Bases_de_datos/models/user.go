package models

import "time"

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Name      string  `gorm:"type:varchar(100);not null"`
	Email     string  `gorm:"uniqueIndex;type:varchar(100);not null"`
	Orders    []Order `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
