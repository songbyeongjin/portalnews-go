package service_clean

import (
	"portal_news/domain_clean/model"
)

type UserService interface {
	UserExistCheck(user *model.User) bool
	CreateUser(user *model.User)
}