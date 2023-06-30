package main

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(setting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ReadTimeout
	writeTimeout := setting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	maxHeaderBytes := 1 << 20
	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logging.Info("[info] start http server listening %s", endPoint)

	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
