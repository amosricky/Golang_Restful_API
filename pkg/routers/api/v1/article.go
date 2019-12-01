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

func GetArticles(c *gin.Context) {

	responseBody := util.NewResponseBody(util.NewBaseError(http.StatusOK, ""))

	for{
		res, err := models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize)
		if err != nil{
			logrus.Errorf("GetArticles :%v", err)
			responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
		}else {
			responseBody.Result = res
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

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
				logrus.Errorf("GetArticle :%v", err)
				responseBody.SetExtendError(util.NewBaseError(http.StatusBadRequest, err.Error()))
			}else {
				responseBody.Result = res
			}
		}
		break
	}

	c.JSON(responseBody.StatusCode(), responseBody)
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(tagId) {
			data := make(map[string]interface {})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logrus.Error(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

//func EditArticle(c *gin.Context) {
//	valid := validation.Validation{}
//
//	id := com.StrTo(c.Param("id")).MustInt()
//	tagId := com.StrTo(c.Query("tag_id")).MustInt()
//	title := c.Query("title")
//	desc := c.Query("desc")
//	content := c.Query("content")
//	modifiedBy := c.Query("modified_by")
//
//	var state int = -1
//	if arg := c.Query("state"); arg != "" {
//		state = com.StrTo(arg).MustInt()
//		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//	}
//
//	valid.Min(id, 1, "id").Message("ID必须大于0")
//	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
//	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
//	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
//	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
//	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
//
//	code := e.INVALID_PARAMS
//	if ! valid.HasErrors() {
//		if ok := models.ExistArticleByID(id); ok{
//			if models.ExistAuthorByID(tagId) {
//				data := make(map[string]interface {})
//				if tagId > 0 {
//					data["tag_id"] = tagId
//				}
//				if title != "" {
//					data["title"] = title
//				}
//				if desc != "" {
//					data["desc"] = desc
//				}
//				if content != "" {
//					data["content"] = content
//				}
//
//				data["modified_by"] = modifiedBy
//
//				models.EditArticle(id, data)
//				code = e.SUCCESS
//			} else {
//				code = e.ERROR_NOT_EXIST_TAG
//			}
//		} else {
//			code = e.ERROR_NOT_EXIST_ARTICLE
//		}
//	} else {
//		for _, err := range valid.Errors {
//			logrus.Error(err)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code" : code,
//		"msg" : e.GetMsg(code),
//		"data" : make(map[string]string),
//	})
//}
//
//func DeleteArticle(c *gin.Context) {
//	id := com.StrTo(c.Param("id")).MustInt()
//
//	valid := validation.Validation{}
//	valid.Min(id, 1, "id").Message("ID必须大于0")
//
//	code := e.INVALID_PARAMS
//	if ! valid.HasErrors() {
//		if ok := models.ExistArticleByID(id); ok{
//			models.DeleteArticle(id)
//			code = e.SUCCESS
//		} else {
//			code = e.ERROR_NOT_EXIST_ARTICLE
//		}
//	} else {
//		for _, err := range valid.Errors {
//			logrus.Error(err)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code" : code,
//		"msg" : e.GetMsg(code),
//		"data" : make(map[string]string),
//	})
//}