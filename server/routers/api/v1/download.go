package v1

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/gin_vue/pkg/app"
	"github.com/JiHanHuang/gin_vue/pkg/e"
	"github.com/JiHanHuang/gin_vue/service/download_service"
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
// @Param name body string false "Name"
// @Param file body string false "BTFile"
// @Param id body int false "ID"
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

type DownloadForm struct {
	ID           int    `form:"id" valid:"Range(0,1)"`
	Name         string `form:"name" valid:"Required;MaxSize(100)"`
	Addr         string `form:"addr" valid:"Required"`
	DownloadPath string `form:"down_path" valid:"Required"`
}

// @Tags New
// @Summary DownloadForm
// @Produce  json
// @Param download body DownloadForm false "Download"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/download [post]
func Download(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form DownloadForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	/*
		err := c.ShouldBindJSON(&form)
		if err != nil {
			appG.Response(http.StatusBadRequest, e.ERROR, nil)
			return
		}
	*/

	if strings.Contains(form.Addr, "thunder://") {
		split := strings.SplitN(form.Addr, "thunder://", 2)
		if split[1] != "" {
			decodeBytes, err := base64.StdEncoding.DecodeString(split[1])
			if err != nil {
				appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
				return
			}

			//remove AA and ZZ
			addr := string(decodeBytes[2 : len(decodeBytes)-2])
			downFile := form.Name
			if i := strings.LastIndex(addr, "/"); i >= 0 {
				downFile = addr[i:]
			}
			ftp := &download_service.FTP{
				Addr:     addr,
				DownPath: form.DownloadPath + "/" + downFile,
			}
			if err := ftp.Download(); err != nil {
				appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
				return
			}
			appG.Response(http.StatusOK, e.SUCCESS, nil)
			return
		}
		appG.Response(http.StatusBadRequest, e.ERROR, "invaild input")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
