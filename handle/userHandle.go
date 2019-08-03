package handle

import (
	"gindemo/model"
	"gindemo/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err ->", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		var result model.User
		model.Db.Table("users").Select("name,email").
			Where("email = ?", user.Email).Limit(1).Scan(&result)
		if result.Email != "" {
			context.String(http.StatusBadRequest, "email已经注册")
		} else {
			result := model.User{Email: user.Email, Password: user.Password, Name: user.Email, RegTime: time.Now().Unix()}
			model.Db.Table("users").Create(&result)
			context.Redirect(http.StatusMovedPermanently, "/")
		}
	}
}

func UserLogin(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")
	if email == "" || password == "" {
		context.String(http.StatusBadRequest, "Email和密码不能为空")
	} else {
		var result model.User
		row := model.Db.Table("users").Where("email = ? AND password=?", email, password).First(&result)
		if row.RecordNotFound() {
			context.String(http.StatusBadRequest, "email或密码不正确")
		} else {
			//登录成功
			row.Scan(&result)
			se, err := util.AesEncrypt(result.Email + "|" + strconv.Itoa(result.ID) + "|" + strconv.FormatInt(time.Now().Unix(), 10))
			if err == nil {
				context.SetCookie("user", se, 3600, "/", "localhost", false, true)
				context.Redirect(http.StatusMovedPermanently, "/")
			}
		}

	}
}

func UserLogout(context *gin.Context) {
	context.SetCookie("user", "", 3600, "/", "localhost", false, true)
	context.Redirect(http.StatusMovedPermanently, "/")
}
