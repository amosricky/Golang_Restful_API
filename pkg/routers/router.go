package routers

import (
	"Golang_Restful_API/pkg/routers/api/v1"
	"Golang_Restful_API/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/author", v1.GetAuthor)
	apiv1.POST("/author", v1.PostAuthor)
	apiv1.PUT("/author/:id", v1.PutAuthor)
	apiv1.DELETE("/author/:id", v1.DeleteAuthor)

	return r
}