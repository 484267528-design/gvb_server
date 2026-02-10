package models

type TagModel struct {
	MODEL
	Title    string         `json:"title"`                          ///标签名称
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"-"` //关联该标签的文章列表，建立article_model_tag表
}
