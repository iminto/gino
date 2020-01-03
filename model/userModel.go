package model

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
}
