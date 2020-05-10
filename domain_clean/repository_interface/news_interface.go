package repository


import (
	"portal_news/domain_clean/model"
)

type NewsRepository interface {
	FindFirstByUrl(url string) *model.News

	FindByPortalsAndTitleLike(portal []string, title string) *[]model.News
	FindByPortalsAndContentLike(portal []string, content string) *[]model.News
	FindByPortalsAndJaTitleLike(portal []string, jaTitle string) *[]model.News
	FindByPortalsAndJaContentLike(portal []string, jaContent string) *[]model.News
}
