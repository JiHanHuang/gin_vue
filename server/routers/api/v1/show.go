package v1

import (
	"net/http"
	"github.com/JiHanHuang/gin_vue/pkg/setting"
	"github.com/JiHanHuang/gin_vue/service/walkdir_service"
	"github.com/gin-gonic/gin"
	"github.com/JiHanHuang/gin_vue/pkg/app"
	"github.com/JiHanHuang/gin_vue/pkg/e"

)

// @Tags New
// @Summary walkdir
// @Param path query string true "./"
// @Router /api/v1/show/files [get]
func ShowFiles(c *gin.Context) {
	appG := app.Gin{C: c}
	realpath := "./" + setting.AppSetting.DownloadSavePath + c.Query("path")
	files := walkdir_service.WalkDir(realpath)
	appG.Response(http.StatusOK, e.SUCCESS, files)
}