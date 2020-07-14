package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/JiHanHuang/stub/docs/swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/JiHanHuang/stub/routers/api/set"
	v1 "github.com/JiHanHuang/stub/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<h1>Welcome Service Stub</h1>`)
	})

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		//test
		apiv1.POST("/post", v1.Tpost)
		apiv1.GET("/get", v1.Tget)
		apiv1.GET("/geturl/*any", v1.TgetUrl)
		apiv1.POST("/posturl/*any", v1.TpostUrl)
		apiv1.GET("/download2", v1.DownFile2)
		apiv1.GET("/download/*any", v1.DownFile)
		apiv1.POST("/upload/", v1.UpFile)
	}
	apiSet := r.Group("/api/set")
	{
		apiSet.POST("/response", set.SetResponse)
	}

	return r
}
