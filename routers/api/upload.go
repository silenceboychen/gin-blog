package api

import (
	"gin-blog/pkg/err"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	code := err.SUCCESS
	data := make(map[string]string)

	file, image, e := c.Request.FormFile("images")
	if e != nil {
		logging.Warn(e)
		code = err.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = err.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = err.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			e := upload.CheckImage(fullPath)
			if e != nil {
				logging.Warn(e)
				code = err.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if e := c.SaveUploadedFile(image, src); e != nil {
				logging.Warn(e)
				code = err.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  err.GetMsg(code),
		"data": data,
	})
}
