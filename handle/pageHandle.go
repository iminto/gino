package handle

import (
	"fmt"
	"gindemo/model"
	"gindemo/service"
	"gindemo/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func Index(ctx *gin.Context) {
	cookie, err := ctx.Cookie("user")
	var user string = ""
	if err == nil {
		text, error := util.AesDecrypt(cookie)
		if error == nil {
			deCookie := strings.Split(text, "|")
			user = deCookie[0]
		} else {
			println(error.Error())
		}
	} else {
		println(err.Error())
	}
	var result []model.User
	result = service.UserList()
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello guy， " + user,
		"userList":result,
	})
}

func Pong(ctx *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}
	ctx.JSON(http.StatusOK, data)
}

func UploadImg(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)
	ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	ctx.SaveUploadedFile(file, "/tmp/"+file.Filename)
}

func JsonInput(ctx *gin.Context) {
	var json interface{}
	if ctx.BindJSON(&json) == nil {
		m := json.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
		ctx.JSON(http.StatusOK, m)
	}

}
