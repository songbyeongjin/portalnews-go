package service

import "time"

type DisplayReviewTemplate struct {
	Portal      string
	NewsUrl     string
	NewsTitle   string
	NewsContent string
	Title       string
	Content     string
	Date        time.Time
}


type MyPageService interface {
	GetReviewByUserId(userId string) *[]DisplayReviewTemplate
}

