package service

import (
	"portal_news/db"
	"portal_news/model"
)

type CreateReviewTemplate struct{
	NewsUrl string
	NewsTitle string
	Portal string
	Press string
	ReviewTitle   string
	ReviewContent string
}

func GetCreateReviewTemplate(userId string, newsUrl string) (*CreateReviewTemplate, bool){
	createReviewTemplate := &CreateReviewTemplate{}
	review := &model.Review{}

	var news = &model.News{}
	newsNotExist :=  db.Instance.Where("url = ?", newsUrl).First(news).RecordNotFound()
	if newsNotExist{
		return nil,false
	}

	createReviewTemplate.Portal = news.Portal
	createReviewTemplate.NewsUrl = newsUrl
	if news.Portal == "naver"{
		createReviewTemplate.NewsTitle = news.TitleJapanese
		createReviewTemplate.Press = news.PressJapanese
	}else{
		createReviewTemplate.NewsTitle = news.Title
		createReviewTemplate.Press = news.Press
	}

	reviewNotExist := db.Instance.Where("news_url = ? AND user_id = ?", newsUrl, userId).First(review).RecordNotFound()
	var modifyMode bool
	if !reviewNotExist {
		createReviewTemplate.ReviewTitle = review.Title
		createReviewTemplate.ReviewContent = review.Content
		modifyMode = true
	}

	return createReviewTemplate, modifyMode
}