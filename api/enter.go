package api

import (
	"gvb_server/api/advert"
	"gvb_server/api/article"
	"gvb_server/api/big_model"
	"gvb_server/api/chat"
	"gvb_server/api/comment"
	"gvb_server/api/data"
	"gvb_server/api/feedback"
	"gvb_server/api/gaode"
	"gvb_server/api/image"
	"gvb_server/api/log"
	"gvb_server/api/log_v2"
	"gvb_server/api/menu"
	"gvb_server/api/message"
	"gvb_server/api/news"
	"gvb_server/api/role"
	"gvb_server/api/system"
	"gvb_server/api/tag"
	"gvb_server/api/user"
)

type ApiGroup struct {
	SystemApiGroup   system.SystemApi
	ImageApiGroup    image.ImageApi
	AdvertApiGroup   advert.AdvertApi
	MenuApiGroup     menu.MenuApi
	UserApiGroup     user.UserApi
	TagApiGroup      tag.TagApi
	MessageApiGroup  message.MessageApi
	ArticleApiGroup  article.ArticleApi
	CommentApiGroup  comment.CommentApi
	NewsApiGroup     news.NewsApi
	ChatApiGroup     chat.ChatApi
	LogApiGroup      log.LogApi
	DataApiGroup     data.DataApi
	LogV2ApiGroup    log_v2.LogApi
	RoleApiGroup     role.RoleApi
	GaodeApiGroup    gaode.GaodeApi
	FeedbackApiGroup feedback.FeedbackApi
	BigModelApiGroup big_model.BigModelApi
}

var ApiGroupApp = new(ApiGroup) // 创建一个ApiGroupApp实例
