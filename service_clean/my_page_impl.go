package service_clean

import (
	"portal_news/domain_clean/repository_interface"
)

type myPageService struct {
	reviewRepository repository_interface.ReviewRepository
	newsRepository repository_interface.NewsRepository
}

func NewMyPageService(reviewRepository repository_interface.ReviewRepository, newsRepository repository_interface.NewsRepository) MyPageService {
	myPageService := myPageService{
		reviewRepository : reviewRepository,
		newsRepository: newsRepository}
	return &myPageService
}

func (m *myPageService) GetReviewByUserId(userId string) *[]DisplayReviewTemplate{
	reviews := m.reviewRepository.FindByUserIdOrderByDateDESC(userId)

	reviewTemplates := &[]DisplayReviewTemplate{}

	for _, r := range *reviews {
		reviewTemplate := DisplayReviewTemplate{}

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