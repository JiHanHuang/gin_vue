package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/gin_vue/pkg/app"
	"github.com/JiHanHuang/gin_vue/pkg/e"
	"github.com/JiHanHuang/gin_vue/service/torrent_service"
)

type TorrentDownloadForm struct {
	ID     int    `form:"id" valid:"Range(0,1)"`
	Name   string `form:"name" valid:"Required;MaxSize(100)"`
	BTFile string `form:"file" valid:"Required"`
}

// @Tags New
// @Summary TorrentDownload
// @Produce  json
// @Param name query string false "Name"
// @Param file query string false "BTFile"
// @Param id query int false "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/download/torrent [post]
func TorrentDownload(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TorrentDownloadForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	torrent := &torrent_service.Torrent{
		ID:       form.ID,
		Name:     form.Name,
		Addr:     "",
		BTFile:   form.BTFile,
		DownPath: "~/download",
	}
	if err := torrent.Download(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
