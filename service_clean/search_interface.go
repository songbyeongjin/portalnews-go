package service_clean

import (
	"net/url"
	"portal_news/domain_clean/model"
)

type SearchService interface {
	GetSearchNews(values url.Values) (*[]model.News, string)
}