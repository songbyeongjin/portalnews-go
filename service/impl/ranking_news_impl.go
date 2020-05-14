package impl

import (
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/service"
)

type rankingNewsService struct {
	rankingNewsRepository repository_interface.RankingNewsRepository
}

func NewRankingNewsService(rankingNewsRepository repository_interface.RankingNewsRepository) service.RankingNewsService {
	rankingNewsService := rankingNewsService{rankingNewsRepository: rankingNewsRepository}

	return &rankingNewsService
}

func (r *rankingNewsService)GetNewsByPortal(portal string) *[]model.RankingNews {
	news := r.rankingNewsRepository.FindByPortal(portal)
	return news
}