package util

import (
	"Golang_Restful_API/pkg/setting"
	"crypto/md5"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetImageName(name string) string {
	extension := path.Ext(name)
	fileName := strings.TrimSuffix(name, extension)
	fileName = EncodeMD5(fileName)

	return fileName + extension
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func GetImagePath() string {
	return setting.ImageSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExtension(fileName string) bool {
	extension := path.Ext(fileName)
	for _, allowExtension := range setting.ImageSetting.ImageAllowExtension {
		if strings.ToUpper(allowExtension) == strings.ToUpper(extension) {
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	content, err := ioutil.ReadAll(f)
	if err != nil {
		logrus.Warn(err)
		return false
	}

	fileSize := len(content)

	return fileSize <= setting.ImageSetting.ImageMaxSize
}

func CheckImageDir(src string) ExtendError {

	var extendError ExtendError

	for{
		dir, err := os.Getwd()
		if err != nil {
			extendError = NewBaseError(http.StatusInternalServerError, err.Error())
			break
		}

		dirPath := dir + "/" + src
		_, err = os.Stat(dirPath)
		isNotExist := os.IsNotExist(err)
		if isNotExist == true{
			err = os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				extendError = NewBaseError(http.StatusInternalServerError, err.Error())
				break
			}
		}
		break
	}
	return extendError
}

func GetImageFullUrl(name string) string {
	return setting.ImageSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}
