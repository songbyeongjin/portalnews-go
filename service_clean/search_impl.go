package service_clean

import (
	"net/url"
	"portal_news/common"
	"portal_news/domain_clean/model"
	"portal_news/domain_clean/repository_interface"
)

type searchService struct {
	newsRepository   repository_interface.NewsRepository
}

func NewSearchService(newsRepository repository_interface.NewsRepository) SearchService {
	searchService := searchService{
		newsRepository: newsRepository}

	return &searchService
}

func (s *searchService) GetSearchNews(values url.Values) (*[]model.News, string){
	var portalStr []string
	var target string
	var language string
	var targetStr string

	if val, ok := values["check-portal"]; ok {
		if val[0] == "all" {
			portalStr = append(portalStr, common.Naver)
			portalStr = append(portalStr, common.Daum)
			portalStr = append(portalStr, common.Nate)
		}else{
			for _, r := range val {
				if r != "all" {
					portalStr = append(portalStr, r)
				}
			}
		}
	}

	if val, ok := values["select-target"]; ok {
		if val[0] == "title" {
			target = "title"
		} else {
			target = "content"
		}
	}

	if val, ok := values["radio-language"]; ok {
		if val[0] == "korean" {
			language = "korean"
		} else {
			language = "japanese"
		}
	}

	if val, ok := values["text-content"]; ok {
		targetStr = val[0]
		if targetStr == "" {
			return nil, ""
		}
	}

	news := &[]model.News{}
	switch language {
	case "korean":
		if target == "title"{
			news = s.newsRepository.FindByPortalsAndTitleLike(portalStr, targetStr)
		}else if target == "content"{
			news = s.newsRepository.FindByPortalsAndContentLike(portalStr, targetStr)
		}
	case "japanese":
		if target == "title"{
			news = s.newsRepository.FindByPortalsAndJaTitleLike(portalStr, targetStr)
		}else if target == "content"{
			news = s.newsRepository.FindByPortalsAndJaContentLike(portalStr, targetStr)
		}
	}

	return news, language
}