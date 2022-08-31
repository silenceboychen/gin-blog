package api

import (
	"gin-blog/pkg/app"
	"gin-blog/pkg/err"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	file, image, e := c.Request.FormFile("images")
	if e != nil {
		logging.Warn(e)
		appG.Response(http.StatusInternalServerError, err.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, err.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, err.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	e = upload.CheckImage(fullPath)
	if e != nil {
		logging.Warn(e)
		appG.Response(http.StatusInternalServerError, err.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}

	if e := c.SaveUploadedFile(image, src); e != nil {
		logging.Warn(e)
		appG.Response(http.StatusInternalServerError, err.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, err.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
