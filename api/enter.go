package api

import (
	"gvb_server/api/advert"
	"gvb_server/api/article"
	"gvb_server/api/chat"
	"gvb_server/api/comment"
	"gvb_server/api/digg"
	"gvb_server/api/image"
	"gvb_server/api/menu"
	"gvb_server/api/message"
	"gvb_server/api/news"
	"gvb_server/api/system"
	"gvb_server/api/tag"
	"gvb_server/api/user"
)

type ApiGroup struct {
	SystemApiGroup  system.SystemApi
	ImageApiGroup   image.ImageApi
	AdvertApiGroup  advert.AdvertApi
	MenuApiGroup    menu.MenuApi
	UserApiGroup    user.UserApi
	TagApiGroup     tag.TagApi
	MessageApiGroup message.MessageApi
	ArticleApiGroup article.ArticleApi
	DiggApiGroup    digg.DiggApi
	CommentApi      comment.CommentApi
	NewsApi         news.NewsApi
	ChatApi         chat.ChatApi
}

var ApiGroupApp = new(ApiGroup) // 创建一个ApiGroupApp实例
