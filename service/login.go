package service

import (
	"portal_news/db"
	"portal_news/model"
)

func UserValidation(user *model.User) bool{
	if user == nil{
		return false
	}

	if len(user.UserId) ==0 || len(user.UserPass) == 0{
		return false
	}

	return true
}


//return exist flag,  already saved exist Obj
func UserExistCheck(user *model.User) (bool,*model.User){
	var userObj = new(model.User)

	return db.Instance.Where("user_id = ?", user.UserId).First(userObj).RecordNotFound(), userObj
}