package v1

import (
	"Golang_Restful_API/pkg/models"
	"Golang_Restful_API/pkg/setting"
	"Golang_Restful_API/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary Get list of authors.
// @Tags author
// @Accept  json
// @Produce  json
// @Success 200 {array} example.GetAuthors
// @Router /author [get]
func GetAuthors(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	name := c.Query("name")
	maps := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	responseBody.Result = models.GetAuthors(util.GetPage(c), setting.AppSetting.PageSize, maps)
	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Get author info by ID.
// @Tags author
// @Accept  json
// @Produce  json
// @Param   id   path   string   true   "Author ID"
// @Success 200 {array} example.GetAuthor
// @Failure 500 {object} example.InvalidID
// @Router /author/{id} [get]
func GetAuthor(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	id := com.StrTo(c.Param("id")).MustInt()

	if id == 0 {
		logrus.Errorf("GetAuthor :%v", util.GetMsg(util.ErrorInvalidID))
		responseBody.SetExtendError(util.NewBaseError(util.ErrorInvalidID, ""))
	}else {
		res := models.GetAuthor(id)
		if res.ID > 0{
			responseBody.Result = res
		}else {
			logrus.Errorf("GetAuthor :%v", util.GetMsg(util.ErrorNotExistAuthor))
			responseBody.SetExtendError(util.NewBaseError(util.ErrorNotExistAuthor, ""))
		}
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Add new author
// @Description </p>[name] : Author name (max:100) </p>[born] : The born year of the author (min:1000, max:2500)
// @Tags author
// @Accept  json
// @Produce  json
// @Param body body example.PostAuthor true " "
// @Success 200 {object} example.OK
// @Router /author [post]
func PostAuthor(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))

	for{
		var addAuthor models.Author
		err := c.BindJSON(&addAuthor)
		if err != nil{
			logrus.Errorf("PostAuthor :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		name := addAuthor.Name
		born := addAuthor.Born

		valid := validation.Validation{}
		valid.Required(name, "name").Message("[name] can't be empty!")
		valid.MaxSize(name, 100, "name").Message("[name] is too long! (max:100)")
		valid.Required(born, "born").Message("[born] can't be empty!")
		valid.Range(born, 1000, 2500, "born").Message("Invalid [born] number! (min:1000, max:2500)")

		if ! valid.HasErrors() {
			models.AddAuthor(name, born)
		}else {
			errMsg := valid.Errors
			logrus.Errorf("PostAuthor :%v", errMsg[0].Message)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, errMsg[0].Message))
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Modify author info
// @Description </p>[name] : Author name (max:100) </p>[born] : The born year of the author (min:1000, max:2500)
// @Tags author
// @Accept  json
// @Produce  json
// @Param   id   path   string   true   "Author ID"
// @Param body body example.PostAuthor true " "
// @Success 200 {object} example.OK
// @Failure 500 {object} example.AuthorNotExist
// @Router /author/{id} [put]
func PutAuthor(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))

	for{
		id := com.StrTo(c.Param("id")).MustInt()
		var editAuthor models.Author
		err := c.BindJSON(&editAuthor)
		if err != nil{
			logrus.Errorf("PutAuthor :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		name := editAuthor.Name
		born := editAuthor.Born

		valid := validation.Validation{}
		valid.Required(name, "name").Message("[name] can't be empty!")
		valid.MaxSize(name, 100, "name").Message("[name] is too long!")
		valid.Required(born, "born").Message("[born] can't be empty!")
		valid.Range(born, 1000, 2500, "born").Message("Invalid [born] number! (min:1000, max:2500)")

		if ! valid.HasErrors() {
			if models.ExistAuthorByID(id) {
				models.EditAuthor(id, name, born)
			} else {
				logrus.Errorf("PutAuthor :%v", util.GetMsg(util.ErrorNotExistAuthor))
				responseBody.SetExtendError(util.NewBaseError(util.ErrorNotExistAuthor, ""))
			}
		}else {
			errMsg := valid.Errors
			logrus.Errorf("PutAuthor :%v", errMsg[0].Message)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, errMsg[0].Message))
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Delete author info
// @Tags author
// @Produce  json
// @Param   id   path   string   true   "Author ID"
// @Success 200 {object} example.OK
// @Failure 500 {object} example.AuthorNotExist
// @Router /author/{id} [delete]
func DeleteAuthor(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	id := com.StrTo(c.Param("id")).MustInt()

	if models.ExistAuthorByID(id) {
		models.DeleteAuthor(id)
	} else {
		logrus.Errorf("DeleteAuthor :%v", util.GetMsg(util.ErrorNotExistAuthor))
		responseBody.SetExtendError(util.NewBaseError(util.ErrorNotExistAuthor, ""))
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}