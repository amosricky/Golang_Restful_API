package v1

import (
	"Golang_Restful_API/pkg/e"
	"Golang_Restful_API/pkg/models"
	"Golang_Restful_API/pkg/setting"
	"Golang_Restful_API/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unknwon/com"
	"net/http"
)


func GetAuthor(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	res := models.GetAuthor(util.GetPage(c), setting.AppSetting.PageSize, maps)

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : res,
	})
}

func PostAuthor(c *gin.Context) {

	var addAuthor models.Author
	err := c.BindJSON(&addAuthor)
	if err != nil{
		logrus.Error(err)
	}

	name := addAuthor.Name
	born := addAuthor.Born

	valid := validation.Validation{}
	valid.Required(name, "name").Message("[Name] can't be empty!")
	valid.MaxSize(name, 100, "name").Message("Name is too long!")
	valid.Required(born, "born").Message("[Born] can't be empty!")
	valid.Range(born, 1000, 2500, "born").Message("Invalid number!")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistAuthorByName(name) {
			code = e.SUCCESS
			models.AddAuthor(name, born)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func PutAuthor(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()
	var editAuthor models.Author
	err := c.BindJSON(&editAuthor)
	if err != nil{
		logrus.Error(err)
	}

	name := editAuthor.Name
	born := editAuthor.Born

	valid := validation.Validation{}
	valid.Required(name, "name").Message("[Name] can't be empty!")
	valid.MaxSize(name, 100, "name").Message("Name is too long!")
	valid.Required(born, "born").Message("[Born] can't be empty!")
	valid.Range(born, 1000, 2500, "born").Message("Invalid number!")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistAuthorByID(id) {
			code = e.SUCCESS
			models.EditAuthor(id, name, born)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func DeleteAuthor(c *gin.Context) {

	id := com.StrTo(c.Param("id")).MustInt()

	code := e.SUCCESS
	if models.ExistAuthorByID(id) {
		models.DeleteAuthor(id)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}