package repository_interface

import(
	"portal_news/domain/model"
)

type UserRepository interface {
	FindFirstByUserId(userId string) *model.User
	Create(user *model.User)
}