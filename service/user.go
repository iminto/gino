package service

import "gindemo/model"

func UserList() []model.User {
	var result []model.User
	model.Db.Table("users").Where("id>?", "0").Find(&result)
	return result
}
