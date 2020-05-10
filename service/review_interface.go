package service

type CreateReviewTemplate struct {
	NewsUrl       string
	NewsTitle     string
	Portal        string
	Press         string
	ReviewTitle   string
	ReviewContent string
}

type JsonBody struct {
	ReviewTitle   string `json:"reviewTitle"`
	ReviewContent string `json:"reviewContent"`
	NewsUrl       string `json:"newsUrl"`
}

type ReviewService interface {
	GetReviewByNewsUrlAndUserId(userID string, newsUrl string) (*CreateReviewTemplate, bool)
	PostReview(jsonBody *JsonBody, userID string)
	UpdateReview(url, userID string, jsonBody *JsonBody) error
	DeleteReview(url, userID string)
}
