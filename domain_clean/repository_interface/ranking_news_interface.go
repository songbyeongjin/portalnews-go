package repository


import (
	"portal_news/domain_clean/model"
)

type RankingNewsRepository interface {
	FindByPortal(portal string) *[]model.Review
}
