package repository_interface

import(
	"portal_news/domain/model"
)

type ReviewRepository interface {
	Create(*model.Review)
	Update(*model.Review, map[string]interface{})
	Delete(*model.Review)
	FindFirstByNewsUrlAndUserId(newsUrl, userID string) *model.Review
	FindByUserIdOrderByDateDESC(userID string) *[]model.Review
}