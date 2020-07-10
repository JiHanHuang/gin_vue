package app

import (
	"net/http"

	"github.com/JiHanHuang/stub/pkg/e"
	"github.com/JiHanHuang/stub/pkg/file"
	"gopkg.in/ini.v1"
)

var responseFileName = "set_response.ini"
var responseFilePath = "runtime/app/"

func SetResponseData(data interface{}, sectionName string) error {

	f, er := file.MustOpen(responseFileName, responseFilePath)
	if er != nil {
		return er
	}
	f.Close()
	fini, err := ini.Load(responseFilePath + responseFileName)
	if err != nil {
		return err
	}
	section := fini.Section(sectionName)
	err = section.ReflectFrom(data)
	if err != nil {
		return err
	}
	err = fini.SaveToIndent(responseFilePath+responseFileName, "\t")
	if err != nil {
		return err
	}
	return nil
}

// Response setting gin.JSON
func (g *Gin) ResponseExt(sectionName string) {
	fini, err := ini.Load(responseFilePath + responseFileName)
	if err != nil {
		g.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
	section, er := fini.GetSection(sectionName)
	if er != nil {
		g.Response(http.StatusOK, e.SUCCESS, nil)
		return
	}
	code := section.Key("Code").MustInt()
	data := section.Key("Data").String()
	ct := section.Key("ContentType").String()
	g.C.Set("ContentType", ct)
	g.C.String(code, data)
	return
}
