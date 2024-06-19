package image

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/response"
	"gvb_server/service/common_service"
)

type ImageListResponse struct {
	models.BannerModel
	BannerCount  int `json:"bannerCount"`  // 关联banner的个数
	ArticleCount int `json:"articleCount"` // 关联文章的个数
}

// ImageListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description 图片列表
// @Param data query models.Page false "查询参数"
// @Router /api/images [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.BannerModel]}
func (ImageApi) ImageListView(c *gin.Context) {
	var cr models.Page
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithCode(gin.ErrorTypeBind, c)
		return
	}

	imageList, count, err := common_service.CommonList(models.BannerModel{}, common_service.Option{
		Page:    cr,
		Likes:   []string{"name"},
		Preload: []string{"MenusBanner"},
		Debug:   true, // 开启debug模式
	})
	var imageIDList []interface{}
	for _, model := range imageList {
		imageIDList = append(imageIDList, model.ID)
	}
	res1, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewTermsQuery("banner_id", imageIDList...)).
		Size(10000).
		Do(context.Background())
	if err != nil {
		return
	}
	var imageIDArticleMap = map[uint]int{}
	for _, hit := range res1.Hits.Hits {
		var model models.ArticleModel
		err = json.Unmarshal(hit.Source, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		val, ok := imageIDArticleMap[model.BannerID]
		if !ok {
			imageIDArticleMap[model.BannerID] = 1
		} else {
			imageIDArticleMap[model.BannerID] = val + 1
		}
	}

	// 算图片id在文章那边有多少个
	var list = make([]ImageListResponse, 0)
	for _, model := range imageList {
		list = append(list, ImageListResponse{
			BannerModel:  model,
			BannerCount:  len(model.MenusBanner),
			ArticleCount: imageIDArticleMap[model.ID],
		})
	}

	response.OkWithList(list, count, c)
}
