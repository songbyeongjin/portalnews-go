package service

import (
	"portal_news/domain/model"
)

type UserService interface {
	UserExistCheck(user *model.User) bool
	CreateUser(user *model.User)
}