package models

// CommentModel 评论表
type CommentModel struct {
	MODEL       `json:"select(c)"`
	SubComments []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(c)"` // 子评论列表

	// --------子评论 belongs to 父评论--------
	ParentCommentID    *uint         `gorm:"comment:父评论ID" json:"parent_comment_id,select(c)"` // 父评论ID（可以为nil）
	ParentCommentModel *CommentModel `gorm:"foreignKey:ParentCommentID" json:"comment_model"`  // 父评论（可以为nil）

	Content      string `gorm:"size:256;comment:评论内容" json:"content,select(c)"`               // 评论内容
	DiggCount    int    `gorm:"size:8;default:0;comment:评论点赞量" json:"digg_count,select(c)"`   // 评论点赞量
	CommentCount int    `gorm:"size:8;default:0;comment:子评论量" json:"comment_count,select(c)"` // 子评论量

	// --------文章 一对多 评论--------
	ArticleID string `gorm:"size:32;comment:评论文章ID" json:"article_id,select(c)"` // 评论文章ID

	// --------评论 belongs to 用户--------
	UserID uint      `gorm:"comment:评论用户ID" json:"user_id,select(c)"` // 评论关联的用户ID
	User   UserModel `gorm:"comment:评论用户" json:"user,select(c)"`      // 评论关联的用户
}
