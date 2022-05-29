package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jamesluo111/gin-blog/middleware/jwt"
	"github.com/jamesluo111/gin-blog/pkg/export"
	"github.com/jamesluo111/gin-blog/pkg/qrcode"
	"github.com/jamesluo111/gin-blog/pkg/setting"
	"github.com/jamesluo111/gin-blog/pkg/upload"
	"github.com/jamesluo111/gin-blog/routers/api"
	v1 "github.com/jamesluo111/gin-blog/routers/api/v1"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	_ "github.com/jamesluo111/gin-blog/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/auth", api.GetAuth)

	r.POST("/upload", api.UploadImage)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apivi := r.Group("/api/v1")
	apivi.Use(jwt.JWT())
	{
		//标签模块
		apivi.GET("/tags", v1.GetTag)
		apivi.POST("/tags", v1.AddTag)
		apivi.PUT("/tags/:id", v1.EditTag)
		apivi.DELETE("/tags/:id", v1.DeleteTag)
		apivi.POST("/tags/export", v1.ExportTag)
		apivi.POST("/tags/import", v1.ImportTag)

		//文章模块
		apivi.GET("/articles", v1.GetArticles)
		apivi.GET("/articles/:id", v1.GetArticle)
		apivi.POST("/articles", v1.AddArticle)
		apivi.PUT("/articles/:id", v1.EditArticle)
		apivi.DELETE("/articles/:id", v1.DeleteArticle)
		apivi.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
