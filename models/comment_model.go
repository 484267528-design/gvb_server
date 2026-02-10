package models

type CommentModel struct {
	MODEL
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`         //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"` //父级评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                                      //父评论ID
	Content            string          `json:"content"`                                                //评论内容
	DiggCount          int             `gorm:"size:8;default:0 "json:"digg_count"`                     //点赞量
	CommentCount       int             `gorm:"size:8;default:0 "json:"comment_count"`                  //评论数
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"article"`                    //关联的文章
	ArticleID          uint            `json:"article_id"`                                             //文章ID
	User               UserModel       `gorm:"foreignKey:UserID" json:"user"`                          //关联的用户
	UserID             uint            `json:"user_id"`                                                //关联的用户ID
}
