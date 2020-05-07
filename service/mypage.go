package service

import (
	"portal_news/db"
	"portal_news/model"
	"time"
)

type DisplayReviewTemplate struct {
	Portal      string
	NewsUrl     string
	NewsTitle   string
	NewsContent string
	Title       string
	Content     string
	Date        time.Time
}

func GetReviewTemplates(userId string) *[]DisplayReviewTemplate {
	reviews := &[]model.Review{}
	db.Instance.Order("date DESC", true).Find(reviews, "user_id=?", userId)

	reviewTemplates := &[]DisplayReviewTemplate{}

	for _, r := range *reviews {
		reviewTemplate := DisplayReviewTemplate{}

		reviewTemplate.NewsUrl = r.NewsUrl
		reviewTemplate.Title = r.Title
		reviewTemplate.Content = r.Content
		reviewTemplate.Date = r.Date

		news := &model.News{}
		db.Instance.Find(news, "url=?", r.NewsUrl)

		if news.Portal == "naver" {
			reviewTemplate.NewsTitle = news.TitleJapanese
			reviewTemplate.NewsContent = news.ContentJapanese
		} else {
			reviewTemplate.NewsTitle = news.Title
			reviewTemplate.NewsContent = news.Content
		}

		reviewTemplate.Portal = news.Portal

		*reviewTemplates = append(*reviewTemplates, reviewTemplate)
	}

	return reviewTemplates
}
