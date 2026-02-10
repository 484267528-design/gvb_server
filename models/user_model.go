package models

import "gvb_server/models/ctype"

type UserModel struct {
	MODEL
	NickName      string           `gorm:"size:36" json:"nick_name" `                               //昵称
	UserName      string           `gorm:"size:36" json:"user_name" `                               //用户名
	Password      string           `gorm:"size:255" json:"password" `                               //密码
	Avatar        string           `gorm:"size:255" json:"avatar_id" `                              //头像ID
	Email         string           `gorm:"size:255" json:"email" `                                  //邮箱
	Tel           string           `gorm:"size:255" json:"tel" `                                    //电话
	Addr          string           `gorm:"size:255" json:"addr" `                                   //地址
	Token         string           `gorm:"size:255" json:"token" `                                  //唯一ID
	IP            string           `gorm:"size:255" json:"ip" `                                     //IP地址
	Role          ctype.Role       `gorm:"size:255;default:1" json:"role" `                         //角色
	SignStatus    ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status" `                    //注册来源
	ArticleModels []ArticleModel   `gorm:"foreignKey:UserID"`                                       //发布的文章列表
	CollectModels []ArticleModel   `gorm:"many2many:collect_article_models;" json:"collect_models"` //收藏的文章列表
}

//User 1 ---- n Article        （一个用户写多篇文章）
//Article 1 ---- n Comment    （一篇文章有多条评论）
//Comment 1 ---- n Comment    （评论支持楼中楼）
//Article n ---- m Tag        （文章 ↔ 标签 多对多）
//User n ---- m Article       （用户 ↔ 收藏文章 多对多）
//Article 1 ---- 1 Banner     （文章有一个封面）
