package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
)

func Daumget(c *gin.Context){
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", "daum").Find(&rankingNews)

	c.HTML(http.StatusOK, "news",gin.H{
		const_val.News:       rankingNews,
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}
