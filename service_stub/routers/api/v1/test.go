package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/app"
	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/file"
)

var filesPath = "./runtime/files/"

// @Tags Test
// @Summary 获取数据
// @Produce  json
// @Param name query string false "自定义返回(可选)"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/get [get]
func Tget(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	if name == "" {
		appG.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
	appG.ResponseExt(name)
}

// @Tags Test
// @Summary 上传数据
// @Produce  json
// @Param post body string false "post" default({"data":"helllo"})
// @Param name query string false "自定义返回(可选)"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/post [post]
func Tpost(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	if name == "" {
		appG.Response(http.StatusOK, e.SUCCESS, string(body))
		return
	}
	appG.ResponseExt(name)
}

// @Tags Test
// @Summary get url信息获取
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/geturl [get]
func TgetUrl(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make([]string, 0, 4)
	body, _ := ioutil.ReadAll(c.Request.Body)
	for k, v := range c.Request.Header {
		data = append(data, fmt.Sprintf("%s:%v", k, v))
	}
	data = append(data, "body:"+string(body))
	data = append(data, "content_len:"+strconv.FormatInt(c.Request.ContentLength, 10))
	appG.Response(http.StatusOK, e.SUCCESS, &data)
}

// @Tags Test
// @Summary post url信息获取
// @Produce  json
// @Param data body string false "Data" default({"data":"helllo"})
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/posturl [post]
func TpostUrl(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make([]string, 0, 4)
	body, _ := ioutil.ReadAll(c.Request.Body)
	for k, v := range c.Request.Header {
		data = append(data, fmt.Sprintf("%s:%v", k, v))
	}
	if !isJSON(string(body)) {
		appG.Response(http.StatusBadRequest, e.ERROR_INVALID_JSON, nil)
		return
	}
	data = append(data, "body:"+string(body))
	data = append(data, "content_len:"+strconv.FormatInt(c.Request.ContentLength, 10))
	appG.Response(http.StatusOK, e.SUCCESS, &data)
}

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

// @Tags Test
// @Summary 下载文件2
// @Param filename query string true "file name"
// @Success 200 {object} gin.Context
// @Router /api/v1/download2 [get]
func DownFile2(c *gin.Context) {
	filename := c.DefaultQuery("filename", "")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filesPath + filename)
}

// @Tags Test
// @Summary 下载文件
// @Param filename query string true "file name"
// @Success 200 {object} octet-stream
// @Router /api/v1/download [get]
func DownFile(c *gin.Context) {
	appG := app.Gin{C: c}
	filename := c.Query("filename")
	if filename == "" {
		fstr := strings.SplitN(c.Request.RequestURI, "/api/v1/download", 2)
		if len(fstr) < 2 || fstr[1] == "/" {
			appG.Response(http.StatusBadRequest, e.ERROR_INVALID_JSON, nil)
			return
		}
		filename = fstr[1][1:]
	}
	file, err := os.Open(filesPath + filename) //Create a file
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	_, err = io.Copy(c.Writer, file)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
}

// @Tags Test
// @Summary 上传文件
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/upload [post]
func UpFile(c *gin.Context) {
	appG := app.Gin{C: c}

	f, header, err := c.Request.FormFile("file")
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	filename := header.Filename

	if err := file.IsNotExistMkDir(filesPath); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}

	newFile, err := os.Create(filesPath + filename)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, f)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
