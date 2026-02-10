package models

import "time"

// 自定义第三章表，记录用户什么时候收藏了什么文章
type UserCollect struct {
	UserID       uint         `gorm:"primary_key"`
	User         UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `json:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
