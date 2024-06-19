package news

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/models/response"
	"gvb_server/service/redis_service"
	"gvb_server/utils/request"
	"io"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signature"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewsResponse struct {
	Code int                      `json:"code"`
	Data []redis_service.NewsData `json:"data"`
	Msg  string                   `json:"msg"`
}

const (
	newsApi = "https://api.codelife.cc/api/top/list"
	timeout = 2 * time.Second
)

// NewsListView 新闻列表
// @Tags 新闻管理
// @Summary 新闻列表
// @Description 新闻列表
// @Param data body params true "新闻参数"
// @Param version header string true "version"
// @Param User-Agent header string true "User-Agent"
// @Param signaturekey header string true "signaturekey"
// @Router /api/news [post]
// @Produce json
// @Success 200 {object} response.Response{data=[]redis_service.NewsData}
func (NewsApi) NewsListView(c *gin.Context) {
	var cr params
	var headers header

	err := c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	key := fmt.Sprintf("%s_%d", cr.ID, cr.Size)
	newsData, _ := redis_service.GetNews(key)
	if len(newsData) != 0 {
		response.OkWithData(newsData, c)
		return
	}

	httpRes, err := request.Post(newsApi, cr, structs.Map(headers), timeout)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	byteData, err := io.ReadAll(httpRes.Body)
	var newsResponse NewsResponse
	err = json.Unmarshal(byteData, &newsResponse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if newsResponse.Code != 200 {
		response.FailWithMessage(err.Error(), c)
		return
	}

	redis_service.SetNews(key, newsResponse.Data)

	response.OkWithData(newsResponse.Data, c)
}
