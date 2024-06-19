package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

type FeedbackRouter struct {
}

func (FeedbackRouter) InitFeedbackRouter(router *gin.RouterGroup) {
	appGroup := api.ApiGroupApp.FeedbackApiGroup
	router.POST("feedback", appGroup.FeedBackCreateView)
	router.GET("feedback", appGroup.FeedBackListView)
	router.DELETE("feedback", appGroup.FeedBackRemoveView)
}
