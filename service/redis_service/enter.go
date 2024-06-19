package redis_service

const (
	ArticleDiggPrefix    = "article_digg"
	ArticleLookPrefix    = "article_look"
	ArticleCommentPrefix = "article_comment"
	CommentDiggPrefix    = "comment_digg"
)

func NewArticleDiggCount() Count {
	return Count{ArticleDiggPrefix}
}

func NewArticleLookCount() Count {
	return Count{ArticleLookPrefix}
}

func NewArticleCommentCount() Count {
	return Count{ArticleCommentPrefix}
}

func NewCommentDiggCount() Count {
	return Count{CommentDiggPrefix}
}
