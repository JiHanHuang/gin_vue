package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/JiHanHuang/gin_vue/docs/swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/JiHanHuang/gin_vue/middleware/info"
	"github.com/JiHanHuang/gin_vue/pkg/export"
	"github.com/JiHanHuang/gin_vue/pkg/qrcode"
	"github.com/JiHanHuang/gin_vue/pkg/upload"
	"github.com/JiHanHuang/gin_vue/routers/api"
	v1 "github.com/JiHanHuang/gin_vue/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(info.MSG())
	//apiv1.Use(jwt.JWT())
	{

		//download
		apiv1.POST("/download/torrent", v1.TorrentDownload)
		apiv1.POST("/download", v1.Download)
		apiv1.GET("/download/list", v1.GetDownloadList)
		//get file from server
		apiv1.GET("/getfile", v1.GetFile)
		//test
		apiv1.POST("/post", v1.Tpost)
		apiv1.GET("/get", v1.Tget)

	}

	return r
}
