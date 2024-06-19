package flags

import (
	"gvb_server/models"
	"gvb_server/service/es_service/indexs"
)

func ESIndex() {
	indexs.CreateIndex(models.FullTextModel{})
	indexs.CreateIndex(models.ArticleModel{})
}
