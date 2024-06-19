package image_service

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
)

var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{"jpg", "jpeg", "png", "gif", "bmp", "webp", "ico", "svg", "tiff"}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 上传失败的消息
}

// ImageUploadService 图片上传方法
/*func (this *ImageService) ImageUploadService(fileHeader *multipart.FileHeader) (res FileUploadResponse) {
	fileName := fileHeader.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, fileName)
	res.FileName = filePath

	// 判断图片后缀是否在白名单中
	suffix := strings.ToLower(path.Ext(fileName)[1:])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件后缀"
		return
	}

	// 判断图片大小
	size := float64(fileHeader.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过%dMB，当前大小为%.2fMB", global.Config.Upload.Size, size)
		return
	}

	// 获取图片的md5值
	fileObj, err := fileHeader.Open()
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		return
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.MD5(byteData)

	// 判断数据库中是否存在该图片
	var bannerModel models.BannerModel
	count := global.DB.Take(&bannerModel, "hash = ?", imageHash).RowsAffected
	if count > 0 { // 存在
		res.FileName = bannerModel.Path
		res.Msg = "图片已存在"
		return
	}

	fileType := ctype.Local
	res.IsSuccess = true
	res.Msg = "图片上传成功"
	// 上传图片到七牛云
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.IsSuccess = false
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛云成功"
		fileType = ctype.QiNiu
	}
	// 将图片写入数据库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return
}
*/

// CreateService 入库
func (this *ImageService) CreateService(filePath, imageHash, fileName string, imageType ctype.ImageType) {
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: imageType,
	})
}
