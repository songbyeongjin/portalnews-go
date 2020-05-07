package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
)

func NaverGet(c *gin.Context){
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", "naver").Find(&rankingNews)

	c.HTML(http.StatusOK, "news",gin.H{
		const_val.News:       rankingNews,
		const_val.LoginFlag : service.GetLoginFlag(c),
		"language" : "japanese",
	})
}

func DaumGet(c *gin.Context){
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", "daum").Find(&rankingNews)

	c.HTML(http.StatusOK, "news",gin.H{
		const_val.News:       rankingNews,
		const_val.LoginFlag : service.GetLoginFlag(c),
		"language" : "korean",
	})
}

func NateGet(c *gin.Context){
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", "nate").Find(&rankingNews)

	c.HTML(http.StatusOK, "news",gin.H{
		const_val.News:       rankingNews,
		const_val.LoginFlag : service.GetLoginFlag(c),
		"language" : "korean",
	})
}


func NaverLanguageGet(c *gin.Context){
	language := c.Param("language")


	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", "naver").Find(&rankingNews)

	c.HTML(http.StatusOK, "news",gin.H{
		const_val.News:       rankingNews,
		const_val.LoginFlag : service.GetLoginFlag(c),
		"language" : language,
	})
}