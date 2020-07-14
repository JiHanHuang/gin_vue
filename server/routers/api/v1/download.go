package v1

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/JiHanHuang/gin_vue/pkg/download"
	"github.com/JiHanHuang/gin_vue/pkg/logging"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/gin_vue/pkg/app"
	"github.com/JiHanHuang/gin_vue/pkg/e"
	"github.com/JiHanHuang/gin_vue/service/download_service"
	"github.com/JiHanHuang/gin_vue/service/torrent_service"
)

type DownloadListForm struct {
	List []status `json:"list"`
}

type status struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Percent int    `json:"percent"`
	Status  string `json:"status"`
}

// @Tags New
// @Summary DownloadListForm
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/download/list [get]
func GetDownloadList(c *gin.Context) {
	appG := app.Gin{C: c}
	list := download_service.GetDownList()
	var downList DownloadListForm
	for _, l := range list {
		var li status
		li.ID = l.GetID()
		li.Name = l.GetName()
		s := l.GetStatus()
		if s.Percent > 100 {
			li.Percent = 100
		} else {
			li.Percent = s.Percent
		}
		logging.Debug(fmt.Sprintln(s))
		switch s.State {
		case download.Failed:
			li.Status = "wrong"
		case download.Running:
			li.Status = "active"
		case download.Finish:
			li.Status = ""
			li.Percent = 100
		default:
			li.Status = ""
		}
		downList.List = append(downList.List, li)
	}

	appG.Response(http.StatusOK, e.SUCCESS, downList)
}

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
	ID           int    `form:"id" valid:"Required"`
	Name         string `form:"name" valid:"Required;MaxSize(100)"`
	Addr         string `form:"addr" valid:"Required"`
	DownloadPath string `form:"downloadPath" valid:"Required"`
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
			if i := strings.LastIndex(addr, "/"); i >= 0 && i+1 < len(addr) {
				downFile = addr[i+1:]
			}
			ftp := download_service.InitDownload(form.ID, addr, form.DownloadPath, downFile)
			if err := ftp.DownloadFile(); err != nil {
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
