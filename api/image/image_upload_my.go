package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/response"
	"gvb_server/plugins/qiniu"
	"gvb_server/service"
	"gvb_server/service/image_service"
	"gvb_server/utils"
	"path"
	"strings"
	"time"
)

// ImageUploadViewMy 上传多个图片，返回图片的url
// @Tags 图片管理
// @Summary 上传多个图片，返回图片的url
// @Description 上传多个图片，返回图片的url
// @Param token header string  true  "token"
// @Accept multipart/form-data
// @Param images formData file true "文件上传"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (ImageApi) ImageUploadViewMy(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fileHeaderList, ok := form.File["images"]
	if !ok {
		response.FailWithMessage("图片不存在", c)
		return
	}

	var resList []image_service.FileUploadResponse

	for _, fileHeader := range fileHeaderList {
		fileName := fileHeader.Filename

		res := image_service.FileUploadResponse{
			FileName:  fileName,
			IsSuccess: false,
			Msg:       "图片上传失败",
		}

		// 判断图片后缀是否在白名单中
		suffix := strings.ToLower(path.Ext(fileName)[1:])
		if !utils.InList(suffix, image_service.WhiteImageList) {
			res.Msg = "非法文件后缀"
			resList = append(resList, res)
			continue
		}

		// 获取图片的md5值
		byteData, imageHash, _ := utils.GetImageMD5(fileHeader)

		// 判断数据库中是否存在该图片
		var bannerModel models.BannerModel
		count := global.DB.Take(&bannerModel, "hash = ?", imageHash).RowsAffected
		if count > 0 { // 存在
			res.Msg = "图片已存在"
			resList = append(resList, res)
			continue
		}

		// 图片大小
		size := float64(fileHeader.Size) / float64(1024*1024)

		// 如果保存到本地
		if !global.Config.QiNiu.Enable {
			// 判断上传路径是否存在
			basePath := global.Config.Upload.Path
			utils.Mkdir(basePath)

			// 判断图片大小
			if size >= float64(global.Config.Upload.Size) {
				res.Msg = fmt.Sprintf("图片大小超过%dMB，当前大小为%.2fMB", global.Config.Upload.Size, size)
				resList = append(resList, res)
				continue
			}
			// 保存图片到本地
			filePath := path.Join(basePath, fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), fileName))
			err = c.SaveUploadedFile(fileHeader, filePath)
			if err != nil {
				global.Log.Error(err)
				resList = append(resList, res)
				continue
			}
			res.IsSuccess = true
			res.Msg = "图片上传成功"
			resList = append(resList, res)

			// 入库
			service.ServiceGroupApp.ImageService.CreateService(filePath, imageHash, fileName, ctype.Local)
		} else {
			// 判断图片大小
			if size >= float64(global.Config.QiNiu.Size) {
				res.Msg = fmt.Sprintf("图片大小超过%dMB，当前大小为%.2fMB", global.Config.QiNiu.Size, size)
				resList = append(resList, res)
				continue
			}
			filePath, err := qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
			if err != nil {
				global.Log.Error(err)
				res.Msg = "图片上传七牛云失败"
				resList = append(resList, res)
				continue
			}
			res.IsSuccess = true
			res.Msg = "图片上传七牛云成功"
			resList = append(resList, res)
			// 入库
			service.ServiceGroupApp.ImageService.CreateService(filePath, imageHash, fileName, ctype.QiNiu)
		}
	}
	response.OkWithData(resList, c)
}
