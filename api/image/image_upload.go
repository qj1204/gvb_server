package image

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/common/response"
	"gvb_server/service"
	"gvb_server/service/image"
	"os"
)

// ImageUploadView 上传图片，返回图片的url
func (this *ImageApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		response.FailWithMessage("图片不存在", c)
		return
	}

	// 判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 不存在就创建
		err = os.MkdirAll(basePath, os.ModePerm) // 递归创建
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []image.FileUploadResponse

	for _, file := range fileList {
		serviceRes := service.ServiceGroupApp.ImageService.ImageUploadService(file)
		// 如果前面上传失败
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 不上传到七牛云，就保存到本地
		if !global.Config.QiNiu.Enable {
			// 保存图片到本地
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	response.OkWithData(resList, c)
}
