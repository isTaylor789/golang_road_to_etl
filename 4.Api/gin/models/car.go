package models

import (
	"time"
)

type Car struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Items     []Item    `gorm:"foreignKey:CarID" json:"items"`
	CreatedAt time.Time `json:"created_at"`
}
