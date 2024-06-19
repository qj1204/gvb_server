package flag

import (
	"gvb_server/models"
	"gvb_server/service/es_service/index"
)

func ESIndex() {
	index.CreateIndex(models.FullTextModel{})
	index.CreateIndex(models.ArticleModel{})
}
