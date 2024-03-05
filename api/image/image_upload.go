package image

// ImageUploadView 上传图片
/*func (this *ImageApi) ImageUploadView(c *gin.Context) {
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

	// 判断上传路径是否存在
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

	for _, fileHeader := range fileHeaderList {
		serviceRes := service.ServiceGroupApp.ImageService.ImageUploadService(fileHeader)
		// 如果某一张图片上传失败，继续上传下一张
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		// 不上传到七牛云，就保存到本地（有问题：这里是先存入数据库再上传，这样不好）
		if !global.Config.QiNiu.Enable {
			// 保存图片到本地
			err = c.SaveUploadedFile(fileHeader, serviceRes.FileName)
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
*/
