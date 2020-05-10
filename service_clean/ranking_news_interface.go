package service_clean

import (
	"portal_news/domain_clean/model"
)

type RankingNewsService interface {
	GetNewsByPortal(portal string) *[]model.RankingNews
}

