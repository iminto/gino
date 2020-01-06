package initRoute

import (
	"gindemo/handle"
	"gindemo/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	route := gin.Default()
	route.Use(middleware.GlobalEcho())
	route.Static("/statics", "./statics")
	route.StaticFile("/favicon.ico", "./favicon.ico")
	route.LoadHTMLGlob("templates/*.tmpl")
	route.MaxMultipartMemory = 8 << 20 // 8 MiB
	route.GET("/", handle.Index)
	userRouter := route.Group("/user")
	{
		userRouter.POST("/register", handle.UserRegister)
		userRouter.POST("/login", handle.UserLogin)
		userRouter.GET("/logout", handle.UserLogout)
		userRouter.GET("/list",  middleware.Auth(),handle.UserList)
	}
	route.GET("/ping", handle.Pong)
	route.POST("/upload", handle.UploadImg) //curl -X POST http://localhost:8090/upload -F "file=@/home/koudai/下载/md5.gif" \-H "Content-Type: multipart/form-data"
	route.POST("/json",handle.JsonInput)
	return route
}
