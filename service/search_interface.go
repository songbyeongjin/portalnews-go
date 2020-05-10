package service

import (
	"net/url"
	"portal_news/domain/model"
)

type SearchService interface {
	GetSearchNews(values url.Values) (*[]model.News, string)
}