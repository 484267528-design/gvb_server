package models

// 用户反馈表
type FeedBackModel struct {
	MODEL
	Email        string `gorm:"size:32" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128" json:"apply_content"`
	IsApply      bool   ` json:"is_apply"`
}
