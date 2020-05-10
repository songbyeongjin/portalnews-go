package repository

import (
	"portal_news/domain_clean/model"
	"portal_news/infra_clean/db"
)

type RankingNewsRepository struct {
	dbHandler db.DbHandler
}

func NewRankingNewsRepository(dbHandler db.DbHandler) RankingNewsRepository {
	rankingNewsRepository := RankingNewsRepository{dbHandler}
	return rankingNewsRepository
}


func (n RankingNewsRepository) FindByPortal(portal string)  *[]model.Review {
	rankingNews := &[]model.Review{}
	n.dbHandler.Conn.Where("portal = ?", portal).Find(rankingNews)

	return rankingNews
}