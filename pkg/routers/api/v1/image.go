package v1

import (
	"Golang_Restful_API/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary Upload image
// @Description  </p>[image] : Accept format (.jpg,.jpeg,.png) Max Size : 5MB
// @Accept  multipart/form-data
// @Produce  json
// @Param image formData file true "Your image file."
// @Success 200 {object} example.UploadImage
// @Success 500 {object} example.ErrorImageFormat
// @Router /upload [post]
func UploadImage(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	data := make(map[string]string)

	for{
		file, image, err := c.Request.FormFile("image")
		if err != nil {
			logrus.Errorf("Upload image :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		if image == nil {
			logrus.Errorf("Image not exist.")
			responseBody.SetExtendError(util.NewBaseError(util.ErrorImageNotExist, ""))
			break
		}

		imageName := util.GetImageName(image.Filename)
		fullPath := util.GetImageFullPath()
		src := fullPath + imageName

		if ! util.CheckImageExtension(imageName) || ! util.CheckImageSize(file) {
			logrus.Errorf("Image file format not correct.")
			responseBody.SetExtendError(util.NewBaseError(util.ErrorImageFormat, ""))
			break
		}

		extendError := util.CheckImageDir(fullPath)
		if extendError != nil {
			logrus.Errorf(extendError.Error())
			responseBody.SetExtendError(extendError)
		} else if err := c.SaveUploadedFile(image, src); err != nil {
			logrus.Errorf("Upload image :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
		} else {
			data["image_url"] = util.GetImageFullUrl(imageName)
			data["encode_file_name"] = imageName
			responseBody.Result = data
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}
