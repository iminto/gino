package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc{
	return func(context *gin.Context) {
		cookie, _ := context.Cookie("user")
		if cookie=="" {
			println("cookie不存在，授权不通过")
			context.AbortWithStatus(500)
		}else{
			println("已经授权")
			context.Next()
		}
	}
}
