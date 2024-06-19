package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/response"
	"gvb_server/service/image_service"
	"gvb_server/utils"
	"gvb_server/utils/jwt"
	"os"
	"path"
	"strings"
	"time"
)

// ImageUploadDataView 上传单个图片，返回图片的url
// @Tags 图片管理
// @Summary 上传单个图片，返回图片的url
// @Description 上上传单个图片，返回图片的url
// @Param token header string  true  "token"
// @Accept multipart/form-data
// @Param image formData file true "文件上传"
// @Router /api/image [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (ImageApi) ImageUploadDataView(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("参数校验失败", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if claims.Role == 3 {
		response.FailWithMessage("游客用户不可上传图片", c)
		return
	}

	fileName := file.Filename
	basePath := global.Config.Upload.Path
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, image_service.WhiteImageList) {
		response.FailWithMessage("非法文件", c)
		return
	}
	// 判断文件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		msg := fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		response.FailWithMessage(msg, c)
		return
	}

	// 创建这个文件夹 /static/uploads/file/nick_name
	dirList, err := os.ReadDir(basePath)
	if err != nil {
		response.FailWithMessage("文件目录不存在", c)
		return
	}
	if !isInDirEntry(dirList, claims.NickName) {
		// 创建这个目录
		err := os.Mkdir(path.Join(basePath, claims.NickName), 077)
		if err != nil {
			global.Log.Error(err)
		}
	}
	// 1.如果存在重名，就加随机字符串 时间戳
	// 2.直接+时间戳
	now := time.Now().Format("20060102150405")
	fileName = nameList[0] + "_" + now + "." + suffix
	filePath := path.Join(basePath, claims.NickName, fileName)

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData("/"+filePath, c)
	return

}

func isInDirEntry(dirList []os.DirEntry, name string) bool {
	for _, entry := range dirList {
		if entry.Name() == name && entry.IsDir() {
			return true
		}
	}
	return false
}
