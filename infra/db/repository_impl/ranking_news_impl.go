package repository_impl

import (
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/infra/db"
)

type RankingNewsRepository struct {
	dbHandler db.Handler
}

func NewRankingNewsRepository(dbHandler db.Handler) repository_interface.RankingNewsRepository {
	rankingNewsRepository := RankingNewsRepository{dbHandler}
	return &rankingNewsRepository
}


func (n *RankingNewsRepository) FindByPortal(portal string)  *[]model.RankingNews {
	rankingNews := &[]model.RankingNews{}
	n.dbHandler.Conn.Where("portal = ?", portal).Find(rankingNews)

	return rankingNews
}