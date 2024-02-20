package models

// CommentModel 评论表
type CommentModel struct {
	MODEL
	SubComments []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"` // 子评论列表

	// 自关联 --------子评论 belongs to 父评论--------
	ParentCommentID    *uint         `gorm:"comment:'父评论ID'" json:"parent_comment_id"`        // 父评论ID
	ParentCommentModel *CommentModel `gorm:"foreignKey:ParentCommentID" json:"comment_model"` // 父评论

	Content      string `gorm:"size:256;comment:'评论内容'" json:"content"`              // 评论内容
	DiggCount    int    `gorm:"size:8;default:0;comment:'评论点赞量'" json:"digg_count"`  // 评论点赞量
	CommentCount int    `gorm:"size:8;default:0;comment:'评论量'" json:"comment_count"` // 子评论量

	// --------文章 一对多 评论--------
	ArticleID string `gorm:"size:32;comment:'评论文章ID'" json:"article_id"` // 评论文章ID

	// --------评论 belongs to 用户--------
	UserID uint      `gorm:"comment:'评论用户ID'" json:"user_id"` // 评论关联的用户ID
	User   UserModel `gorm:"comment:'评论用户'" json:"user"`      // 评论关联的用户
}
