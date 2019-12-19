package routers

import (
	_ "Golang_Restful_API/docs"
	"Golang_Restful_API/pkg/routers/api/v1"
	"Golang_Restful_API/pkg/setting"
	"Golang_Restful_API/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:  []string{"content-type", "authorization"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))
	r.StaticFS("/images", http.Dir(util.GetImageFullPath()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8000/swagger/doc.json")))

	gin.SetMode(setting.ServerSetting.RunMode)

	apiv1 := r.Group("/api/v1")
	apiv1.POST("/register", v1.Register)
	apiv1.POST("/login", v1.Login)

	apiv1.Use(util.JWTAuth())

	apiv1.POST("/upload", v1.UploadImage)

	apiv1.GET("/author", v1.GetAuthors)
	apiv1.GET("/author/:id", v1.GetAuthor)
	apiv1.POST("/author", v1.PostAuthor)
	apiv1.PUT("/author/:id", v1.PutAuthor)
	apiv1.DELETE("/author/:id", v1.DeleteAuthor)

	apiv1.GET("/article", v1.GetArticles)
	apiv1.GET("/article/:id", v1.GetArticle)
	apiv1.POST("/article", v1.AddArticle)
	apiv1.PUT("/article/:id", v1.PutArticle)
	apiv1.DELETE("/article/:id", v1.DeleteArticle)

	return r
}