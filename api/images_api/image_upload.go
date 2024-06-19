package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qiniu"
	"gvb_server/service"
	"gvb_server/service/image_service"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"path"
	"strings"
	"time"
)

// ImageUploadView 上传多个图片，返回图片的url
// @Tags 图片管理
// @Summary 上传多个图片，返回图片的url
// @Description 上传多个图片，返回图片的url
// @Param token header string  true  "token"
// @Accept multipart/form-data
// @Param images formData file true "文件上传"
// @Router /api/images [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageUploadView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	if claims.Role == 3 {
		res.FailWithMessage("游客用户不可上传图片", c)
		return
	}
	// 上传多个图片
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileHeaderList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	var resList []image_service.FileUploadResponse

	for _, fileHeader := range fileHeaderList {
		fileName := fileHeader.Filename
		resp := image_service.FileUploadResponse{
			FileName:  fileName,
			IsSuccess: false,
			Msg:       "图片上传失败",
		}

		// 判断图片后缀是否在白名单中
		suffix := strings.ToLower(path.Ext(fileName)[1:])
		if !utils.InList(suffix, image_service.WhiteImageList) {
			resp.Msg = "非法文件后缀"
			resList = append(resList, resp)
			continue
		}

		// 获取图片的md5值
		byteData, imageHash, _ := utils.GetImageMD5(fileHeader)

		// 判断数据库中是否存在该图片
		var bannerModel models.BannerModel
		count := global.DB.Take(&bannerModel, "hash = ?", imageHash).RowsAffected
		if count > 0 { // 存在
			resp.Msg = "图片已存在"
			resList = append(resList, resp)
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
				resp.Msg = fmt.Sprintf("图片大小超过%dMB，当前大小为%.2fMB", global.Config.Upload.Size, size)
				resList = append(resList, resp)
				continue
			}
			// 保存图片到本地
			filePath := path.Join(basePath, fmt.Sprintf("%s_%s", time.Now().Format("20060102150405"), fileName))
			err = c.SaveUploadedFile(fileHeader, filePath)
			if err != nil {
				global.Log.Error(err)
				resList = append(resList, resp)
				continue
			}
			resp.IsSuccess = true
			resp.Msg = "图片上传成功"
			resList = append(resList, resp)

			// 入库
			service.ServiceApp.ImageService.CreateService(filePath, imageHash, fileName, ctype.Local)
		} else {
			// 判断图片大小
			if size >= global.Config.QiNiu.Size {
				resp.Msg = fmt.Sprintf("图片大小超过%.2fMB，当前大小为%.2fMB", global.Config.QiNiu.Size, size)
				resList = append(resList, resp)
				continue
			}
			filePath, err := qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
			if err != nil {
				global.Log.Error(err)
				resp.Msg = "图片上传七牛云失败"
				resList = append(resList, resp)
				continue
			}
			resp.IsSuccess = true
			resp.Msg = "图片上传七牛云成功"
			resList = append(resList, resp)
			// 入库
			service.ServiceApp.ImageService.CreateService(filePath, imageHash, fileName, ctype.QiNiu)
		}
	}
	res.OkWithData(resList, c)
}
