package redis_service

const (
	ArticleDiggPrefix    = "article_digg"    // 文章浏览量
	ArticleLookPrefix    = "article_look"    // 文章评论数
	ArticleCommentPrefix = "article_comment" // 文章点赞数
	CommentDiggPrefix    = "comment_digg"    // 评论点赞数
)

func NewArticleDiggCount() CountDB {
	return CountDB{
		Index: ArticleDiggPrefix,
	}
}
func NewArticleLookCount() CountDB {
	return CountDB{
		Index: ArticleLookPrefix,
	}
}
func NewArticleCommentCount() CountDB {
	return CountDB{
		Index: ArticleCommentPrefix,
	}
}
func NewCommentDiggCount() CountDB {
	return CountDB{
		Index: CommentDiggPrefix,
	}
}
