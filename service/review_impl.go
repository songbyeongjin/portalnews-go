package service

import (
	"fmt"
	"portal_news/common"
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"time"
)

type reviewService struct {
	reviewRepository repository_interface.ReviewRepository
	newsRepository   repository_interface.NewsRepository
}

func NewReviewService(reviewRepository repository_interface.ReviewRepository, newsRepository repository_interface.NewsRepository) ReviewService {
	reviewService := reviewService{
		reviewRepository: reviewRepository,
		newsRepository: newsRepository}

	return &reviewService
}

func (r *reviewService) GetReviewByNewsUrlAndUserId(userID string, newsUrl string) (*CreateReviewTemplate, bool) {
	createReviewTemplate := &CreateReviewTemplate{}

	news := r.newsRepository.FindFirstByUrl(newsUrl)
	if news == nil {
		return nil, false
	}

	createReviewTemplate.Portal = news.Portal
	createReviewTemplate.NewsUrl = newsUrl
	if news.Portal == common.Naver {
		createReviewTemplate.NewsTitle = news.TitleJapanese
		createReviewTemplate.Press = news.PressJapanese
	} else {
		createReviewTemplate.NewsTitle = news.Title
		createReviewTemplate.Press = news.Press
	}

	review := r.reviewRepository.FindFirstByNewsUrlAndUserId(newsUrl, userID)
	var modifyMode bool
	if review != nil {
		createReviewTemplate.ReviewTitle = review.Title
		createReviewTemplate.ReviewContent = review.Content
		modifyMode = true
	}

	return createReviewTemplate, modifyMode
}

func (r *reviewService) PostReview(jsonBody *JsonBody, userID string){
	review := r.reviewRepository.FindFirstByNewsUrlAndUserId(jsonBody.NewsUrl, userID)
	if review != nil {
		r.reviewRepository.Update(review, map[string]interface{}{"title": jsonBody.ReviewTitle, "content": jsonBody.ReviewContent})
	} else {
		createReview := &model.Review{
			UserId:  userID,
			NewsUrl: jsonBody.NewsUrl,
			Title:   jsonBody.ReviewTitle,
			Content: jsonBody.ReviewContent,
			Date:    time.Now(),
		}
		r.reviewRepository.Create(createReview)
	}
}

func (r *reviewService) UpdateReview(url, userID string, jsonBody *JsonBody) error{
	review := r.reviewRepository.FindFirstByNewsUrlAndUserId(url, userID)

	if review == nil {
		return fmt.Errorf("target url is uncorrect")
	} else {
		r.reviewRepository.Update(review, map[string]interface{}{"title": jsonBody.ReviewTitle, "content": jsonBody.ReviewContent})
	}

	return nil
}


func (r *reviewService) DeleteReview(url, userID string) {
	var review = &model.Review{}
	review = r.reviewRepository.FindFirstByNewsUrlAndUserId(url, userID)

	r.reviewRepository.Delete(review)
}