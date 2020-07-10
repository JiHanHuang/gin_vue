package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JiHanHuang/stub/pkg/gredis"
	"github.com/JiHanHuang/stub/pkg/logging"
	"github.com/JiHanHuang/stub/pkg/setting"
	"github.com/JiHanHuang/stub/pkg/util"
	"github.com/JiHanHuang/stub/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin-VUE API
// @version 1.0
// @description An example of gin+vue
// @termsOfService https://github.com/JiHanHuang/stub

func main() {
	gin.SetMode("debug")

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
