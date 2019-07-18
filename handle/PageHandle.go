package handle

import (
	"fmt"
	"gindemo/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Index(ctx *gin.Context){
	cookie, err := ctx.Cookie("user")
	var user string=""
	if err==nil {
		text,error:=util.AesDecrypt(cookie)
		if error==nil {
			deCookie := strings.Split(text, "|")
			user=deCookie[0]
		}else{
			println(error.Error())
		}
	}else{
		println(err.Error())
	}
	ctx.HTML(http.StatusOK,"index.tmpl",gin.H{
		"title": "hello guy， " + user,
	})
}

func Pong(ctx *gin.Context){
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}
	ctx.JSON(http.StatusOK,data)
}

func UploadImg(ctx *gin.Context){
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	ctx.SaveUploadedFile(file,"/tmp/"+file.Filename)
}