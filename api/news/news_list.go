package news

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/models/common/response"
	"gvb_server/service/redis"
	"gvb_server/utils/request"
	"io"
	"time"
)

type Params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type Header struct {
	Signaturekey string `form:"signaturekey" structs:"signature"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewsResponse struct {
	Code int               `json:"code"`
	Data []models.NewsData `json:"data"`
	Msg  string            `json:"msg"`
}

const (
	newsApi = "https://api.codelife.cc/api/top/list"
	timeout = 2 * time.Second
)

func (api *NewsApi) NewsListView(c *gin.Context) {
	var params Params
	var header Header

	err := c.ShouldBindJSON(&params)
	err = c.ShouldBindHeader(&header)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	key := fmt.Sprintf("%s_%d", params.ID, params.Size)
	newsData, _ := redis.GetNews(key)
	if len(newsData) != 0 {
		response.OkWithData(newsData, c)
		return
	}

	httpRes, err := request.Post(newsApi, params, structs.Map(header), timeout)
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

	redis.SetNews(key, newsResponse.Data)

	response.OkWithData(newsResponse.Data, c)
}
