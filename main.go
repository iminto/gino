package main

import (
	"gindemo/config"
	"gindemo/initRoute"
	"gindemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	model.InitDB(&config.Config)
	//f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := initRoute.SetupRoute()
	s := &http.Server{
		Addr:           config.Config.HostAddr,
		Handler:        router,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	//s.ListenAndServeTLS("gd.pem","gd-key.pem")
}
