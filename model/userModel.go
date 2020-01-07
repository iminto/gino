package model

import "time"

type UserModel struct {
	Email         string `form:"email"  binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}

type User struct {
	ID       int `gorm:"AUTO_INCREMENT"`
	Username string
	Name     string
	Email    string
	Password string
	RegTime  int64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Result struct{
	Code int
	Data interface{}
}
