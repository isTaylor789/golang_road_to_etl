package models

import "time"

type Car struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Items     []Item `gorm:"foreignKey:CarID"`
	CreatedAt time.Time
}

type Item struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CarID     string `gorm:"type:uuid"`
	Name      string
	CreatedAt time.Time
}
