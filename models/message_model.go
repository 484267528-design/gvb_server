package models

type MessageModel struct {
	MODEL
	SendUserId       uint      `gorm:"primaryKey" json:"send_user_id"`
	SendUserModel    UserModel `gorm:"foreignKey:SendUserId" json:"-"`
	SendUserNickname string    `gorm:"size:255" json:"send_user_nickname"`
	SendUserAvatar   string    ` json:"send_user_avatar"`

	RevUserID       uint      `gorm:"primaryKey" json:"rev_user_id"`
	RevUserModel    UserModel `gorm:"foreignKey:RevUserId" json:"-"`
	RevUserNickname string    `gorm:"size:255" json:"rev_user_nickname"`
	RevUserAvatar   string    ` json:"rev_user_avatar"`
	IsRead          bool      `gorm:"default:false" json:"is_read"`
	Content         string    ` json:"content"`
}
