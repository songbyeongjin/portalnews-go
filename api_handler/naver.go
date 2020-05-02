package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/db"
	"portal_news/model"
)

func Naver(c *gin.Context){
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", "naver").Find(&rankingNews)

	//e, _ := json.Marshal(news)

	//c.Header("Content-Type", "application/json")
	c.HTML(http.StatusOK, "news", rankingNews)
}
