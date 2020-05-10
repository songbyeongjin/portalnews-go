package repository

import(
	"portal_news/domain_clean/model"
)

type UserRepository interface {
	FindFirstByUserId(userId string) *model.User
}