package repository_interface


import (
	"portal_news/domain/model"
)

type RankingNewsRepository interface {
	FindByPortal(portal string) *[]model.RankingNews
}
