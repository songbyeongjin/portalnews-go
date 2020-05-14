package impl

import (
	"portal_news/domain/repository_interface"
	"portal_news/service"
)

type myPageService struct {
	reviewRepository repository_interface.ReviewRepository
	newsRepository repository_interface.NewsRepository
}

func NewMyPageService(reviewRepository repository_interface.ReviewRepository, newsRepository repository_interface.NewsRepository) service.MyPageService {
	myPageService := myPageService{
		reviewRepository : reviewRepository,
		newsRepository: newsRepository}
	return &myPageService
}

func (m *myPageService) GetReviewByUserId(userId string) *[]service.DisplayReviewTemplate {
	reviews := m.reviewRepository.FindByUserIdOrderByDateDESC(userId)

	reviewTemplates := &[]service.DisplayReviewTemplate{}

	for _, r := range *reviews {
		reviewTemplate := service.DisplayReviewTemplate{}

		reviewTemplate.NewsUrl = r.NewsUrl
		reviewTemplate.Title = r.Title
		reviewTemplate.Content = r.Content
		reviewTemplate.Date = r.Date

		news := m.newsRepository.FindFirstByUrl(r.NewsUrl)

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