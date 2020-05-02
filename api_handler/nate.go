package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/db"
	"portal_news/model"
)

func Nate(c *gin.Context){
	news := []model.RankingNews{}

	db.Instance.Where("portal = ?", "nate").Find(&news)

	c.HTML(http.StatusOK, "news", news)
}

