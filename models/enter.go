package models

import "time"

type MODEL struct {
	ID        int       `gorm:"primary_key;AUTO_INCREMENT" json:"id" structs:"-"`
	CreatedAt time.Time `json:"created_at" structs:"-"`
	UpdatedAt time.Time `json:"updated_at" structs:"-"`
}
type RemoveResult struct {
	IDList []int `json:"id_list"`
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}
