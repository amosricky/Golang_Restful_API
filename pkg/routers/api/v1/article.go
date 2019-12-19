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

// @Summary Get list of articles.
// @Tags article
// @Accept  json
// @Produce  json
// @Success 200 {array} example.GetArticles
// @Router /article [get]
func GetArticles(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))

	for{
		res, err := models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize)
		if err != nil{
			logrus.Errorf("GetArticles :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
		}else {
			responseBody.Result = res
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Get article info by ID.
// @Tags article
// @Accept  json
// @Produce  json
// @Param   id   path   string   true   "Article ID"
// @Success 200 {array} example.GetArticle
// @Failure 500 {object} example.ArticleNotExist
// @Router /article/{id} [get]
func GetArticle(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))

	for{
		id := com.StrTo(c.Param("id")).MustInt()

		if id==0{
			logrus.Errorf("GetAuthor :%v", util.GetMsg(util.ErrorInvalidID))
			responseBody.SetExtendError(util.NewBaseError(util.ErrorInvalidID, ""))
		}else {
			res, err := models.GetArticle(id)
			if err != nil{
				logrus.Errorf("GetArticle :%v", err.Error())
				responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			}else if res.ID > 0{
				responseBody.Result = res
			}else {
				logrus.Errorf("GetArticle :%v", util.GetMsg(util.ErrorNotExistArticle))
				responseBody.SetExtendError(util.NewBaseError(util.ErrorNotExistArticle, ""))
			}
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Add new article
// @Description </p>[author_id] : Author id (min:1)</p>[title] : The article title </p>[desc] : Description of article </p>[content] : Content of article </p>[imageUrl] : Url of image
// @Tags article
// @Accept  json
// @Produce  json
// @Param body body example.PostArticle true " "
// @Success 200 {object} example.OK
// @Router /article [post]
func AddArticle(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	data := make(map[string]interface{})

	for{
		var addArticle models.Article
		err := c.BindJSON(&addArticle)
		if err != nil{
			logrus.Errorf("PostArticle :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		authorID := addArticle.AuthorID
		title := addArticle.Title
		desc := addArticle.Desc
		content := addArticle.Content
		imageUrl := addArticle.ImageUrl

		data["authorID"] = authorID
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["imageUrl"] = imageUrl

		valid := validation.Validation{}
		valid.Min(authorID, 1, "authorID").Message("Invalid [author_id] (min:1)")
		valid.Required(title, "title").Message("[title] can't be empty!")
		valid.Required(desc, "desc").Message("[desc] can't be empty!")
		valid.Required(content, "content").Message("[content] can't be empty!")
		valid.Required(imageUrl, "imageUrl").Message("[imageUrl] can't be empty!")

		if ! valid.HasErrors() {
			err :=models.AddArticle(data)
			if err != nil{
				logrus.Errorf("PostArticle :%v", err.Error())
				responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			}
		}else {
			errMsg := valid.Errors
			logrus.Errorf("PostArticle :%v", errMsg[0].Message)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, errMsg[0].Message))
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Modify article info
// @Description </p>[author_id] : Author id (min:1)</p>[title] : The article title </p>[desc] : Description of article </p>[content] : Content of article </p>[imageUrl] : Url of image
// @Tags article
// @Accept  json
// @Produce  json
// @Param   id   path   string   true   "Article ID"
// @Param body body example.PostArticle true " "
// @Success 200 {object} example.OK
// @Failure 500 {object} example.ArticleNotExist
// @Router /article/{id} [put]
func PutArticle(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	data := make(map[string]interface{})

	for{
		id := com.StrTo(c.Param("id")).MustInt()
		var editArticle models.Article
		err := c.BindJSON(&editArticle)
		if err != nil{
			logrus.Errorf("PutArticle :%v", err.Error())
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			break
		}

		authorID := editArticle.AuthorID
		title := editArticle.Title
		desc := editArticle.Desc
		content := editArticle.Content
		imageUrl := editArticle.ImageUrl

		data["authorID"] = authorID
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["imageUrl"] = imageUrl

		valid := validation.Validation{}
		valid.Min(authorID, 1, "authorID").Message("Invalid [author_id] (min:1)")
		valid.Required(title, "title").Message("[title] can't be empty!")
		valid.Required(desc, "desc").Message("[desc] can't be empty!")
		valid.Required(content, "content").Message("[content] can't be empty!")
		valid.Required(imageUrl, "imageUrl").Message("[imageUrl] can't be empty!")

		if ! valid.HasErrors() {
			if models.ExistArticleByID(id) {
				err := models.EditArticle(id, data)
				if err != nil{
					logrus.Errorf("PutArticle :%v", err.Error())
					responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
				}
			} else {
				logrus.Errorf("PutArticle :%v", util.GetMsg(util.ErrorNotExistArticle))
				responseBody.SetExtendError(util.NewBaseError(util.ErrorNotExistArticle, ""))
			}
		}else {
			errMsg := valid.Errors
			logrus.Errorf("PutArticle :%v", errMsg[0].Message)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, errMsg[0].Message))
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

// @Summary Delete article info
// @Tags article
// @Produce  json
// @Param   id   path   string   true   "Article ID"
// @Success 200 {object} example.OK
// @Failure 500 {object} example.ArticleNotExist
// @Router /article/{id} [delete]
func DeleteArticle(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))
	id := com.StrTo(c.Param("id")).MustInt()

	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
	} else {
		logrus.Errorf("DeleteArticle :%v", util.GetMsg(util.ErrorNotExistArticle))
		responseBody.SetExtendError(util.NewBaseError(util.ErrorNotExistArticle, ""))
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}