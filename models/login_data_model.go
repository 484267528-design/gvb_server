package models

type LoginDataModel struct {
	MODEL
	UserID    uint      ` json:"user_id"`
	UserModel UserModel `gorm:"foreignkey:UserID" json:"-"`
	Ip        string    `gorm:"size:20" json:"ip"`
	Nickname  string    `gorm:"size:20" json:"nickname"`
	Token     string    `gorm:"size:255" json:"token"`
	Device    string    `gorm:"size:20" json:"device"`
	Addr      string    `gorm:"size:20" json:"addr"`
}
