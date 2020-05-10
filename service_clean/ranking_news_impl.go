package service_clean

import (
	"portal_news/domain_clean/model"
	"portal_news/domain_clean/repository_interface"
)

type rankingNewsService struct {
	rankingNewsRepository repository_interface.RankingNewsRepository
}

func NewRankingNewsService(rankingNewsRepository repository_interface.RankingNewsRepository) RankingNewsService {
	rankingNewsService := rankingNewsService{rankingNewsRepository: rankingNewsRepository}

	return &rankingNewsService
}

func (r *rankingNewsService)GetNewsByPortal(portal string) *[]model.RankingNews {
	news := r.rankingNewsRepository.FindByPortal(portal)
	return news
}