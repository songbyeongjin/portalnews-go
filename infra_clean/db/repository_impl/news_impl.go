package repository

import (
	"portal_news/domain_clean/model"
	"portal_news/infra_clean/db"
)

type NewsRepository struct {
	db.DbHandler
}

func NewNewsRepository(dbHandler db.DbHandler) NewsRepository {
	newsRepository := NewsRepository{dbHandler}
	return newsRepository
}

func (newsRepo *NewsRepository) FindFirstByUrl(url string)  *model.News {
	news := &model.News{}
	newsRepo.DbHandler.Conn.Where("url = ?", url).First(news)

	return news
}

func (newsRepo *NewsRepository) FindByPortalsAndTitleLike(portals []string, title string)  *[]model.News {
	news := &[]model.News{}

	newsRepo.DbHandler.Conn.Where("portal IN (?) AND title" + " LIKE  ?", portals, `%`+title+`%`).Find(news)

	return news
}

func (newsRepo *NewsRepository) FindByPortalsAndContentLike(portals []string, content string)  *[]model.News {
	news := &[]model.News{}

	newsRepo.DbHandler.Conn.Where("portal IN (?) AND content" + " LIKE  ?", portals, `%`+content+`%`).Find(news)

	return news
}

func (newsRepo *NewsRepository) FindByPortalsAndJaTitleLike(portals []string, jaTitle string)  *[]model.News {
	news := &[]model.News{}

	newsRepo.DbHandler.Conn.Where("portal IN (?) AND title_ja" + " LIKE  ?", portals, `%`+jaTitle+`%`).Find(news)

	return news
}

func (newsRepo *NewsRepository) FindByPortalsAndJaContentLike(portals []string, jaContent string)  *[]model.News {
	news := &[]model.News{}

	newsRepo.DbHandler.Conn.Where("portal IN (?) AND content_ja" + " LIKE  ?", portals, `%`+jaContent+`%`).Find(news)

	return news
}