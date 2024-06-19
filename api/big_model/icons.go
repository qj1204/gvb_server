package big_model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gvb_server/models"
	"gvb_server/models/response"
	"os"
	"path"
)

func (BigModelApi) IconsView(c *gin.Context) {
	dir, err := os.ReadDir("uploads/role_icons")
	if err != nil {
		logrus.Error(err)
		response.FailWithMessage("目录不存在", c)
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

	response.OkWithData(list, c)
}
