package models

import "time"

type MODEL struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
