package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/gin_vue/pkg/app"
	"github.com/JiHanHuang/gin_vue/pkg/e"
)

var testPath = "./runtime/test/test.json"

// @Tags Test
// @Summary TorrentDownload
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/get [get]
func Tget(c *gin.Context) {
	appG := app.Gin{C: c}
	b, err := ioutil.ReadFile(testPath)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, string(b))
}

type TpostForm struct {
	ID           int    `form:"id" valid:"Range(0,1)"`
	Name         string `form:"name" valid:"Required;MaxSize(100)"`
	Addr         string `form:"addr" valid:"Required"`
	DownloadPath string `form:"down_path" valid:"Required"`
}

// @Tags Test
// @Summary TpostForm
// @Produce  json
// @Param post body TpostForm false "post"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/post [post]
func Tpost(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form TpostForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	d, err := json.Marshal(form)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	err = ioutil.WriteFile(testPath, d, 0644)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
