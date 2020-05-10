package service

import (
	"portal_news/domain/model"
)

type RankingNewsService interface {
	GetNewsByPortal(portal string) *[]model.RankingNews
}

