package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/db"
	"portal_news/model"
)

func Daum(c *gin.Context){
	news := []model.RankingNews{}

	db.Instance.Where("portal = ?", "daum").Find(&news)

	c.HTML(http.StatusOK, "news", news)
}
