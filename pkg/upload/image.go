package upload

import (
	"Golang_Restful_API/pkg/util"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"Golang_Restful_API/pkg/setting"
)

func GetImageName(name string) string {
	extension := path.Ext(name)
	fileName := strings.TrimSuffix(name, extension)
	fileName = util.EncodeMD5(fileName)

	return fileName + extension
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExtension(fileName string) bool {
	extension := path.Ext(fileName)
	for _, allowExtension := range setting.AppSetting.ImageAllowExtension {
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

	return fileSize <= setting.AppSetting.ImageMaxSize
}

func CheckImageDir(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	dirPath := dir + "/" + src

	_, err = os.Stat(dirPath)
	isNotExist := os.IsNotExist(err)

	if isNotExist == true{
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}