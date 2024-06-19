package big_model_api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/models"
	"gvb_server/models/res"
	"os"
	"path"
)

func (BigModelApi) IconsView(c *gin.Context) {
	dir, err := os.ReadDir("uploads/role_icons")
	if err != nil {
		logrus.Error(err)
		res.FailWithMessage("目录不存在", c)
		return
	}
	var list []models.Options[string]
	for _, entry := range dir {
		key := "/" + path.Join("uploads/role_icons", entry.Name())
		list = append(list, models.Options[string]{
			Label: key,
			Value: key,
		})
	}

	res.OkWithData(list, c)
}
