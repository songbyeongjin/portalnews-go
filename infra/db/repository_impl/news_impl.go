package repository_impl

import (
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/infra/db"
)

type NewsRepository struct {
	dbHandler db.DbHandler
}

func NewNewsRepository(dbHandler db.DbHandler) repository_interface.NewsRepository {
	newsRepository := NewsRepository{dbHandler}
	return &newsRepository
}

func (n *NewsRepository) FindFirstByUrl(url string)  *model.News {
	news := &model.News{}
	n.dbHandler.Conn.Where("url = ?", url).First(news)

	return news
}

func (n *NewsRepository) FindByPortalsAndTitleLike(portals []string, title string)  *[]model.News {
	news := &[]model.News{}

	n.dbHandler.Conn.Where("portal IN (?) AND title" + " LIKE  ?", portals, `%`+title+`%`).Find(news)

	if  len(*news) == 0{
		return nil
	}

	return news
}

func (n *NewsRepository) FindByPortalsAndContentLike(portals []string, content string)  *[]model.News {
	news := &[]model.News{}

	n.dbHandler.Conn.Where("portal IN (?) AND content" + " LIKE  ?", portals, `%`+content+`%`).Find(news)

	if  len(*news) == 0{
		return nil
	}

	return news
}

func (n *NewsRepository) FindByPortalsAndJaTitleLike(portals []string, jaTitle string)  *[]model.News {
	news := &[]model.News{}

	n.dbHandler.Conn.Where("portal IN (?) AND title_ja" + " LIKE  ?", portals, `%`+jaTitle+`%`).Find(news)

	if  len(*news) == 0{
		return nil
	}

	return news
}

func (n *NewsRepository) FindByPortalsAndJaContentLike(portals []string, jaContent string)  *[]model.News {
	news := &[]model.News{}

	n.dbHandler.Conn.Where("portal IN (?) AND content_ja" + " LIKE  ?", portals, `%`+jaContent+`%`).Find(news)

	if  len(*news) == 0{
		return nil
	}

	return news
}

