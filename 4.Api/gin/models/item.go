package models

import (
	"time"
)

type Item struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CarID     string    `gorm:"type:uuid" json:"car_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
