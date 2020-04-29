package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/db"
	"portal_news/model"
)

func Nate(c *gin.Context){
	news := []model.News{}

	db.Instance.Where("portal = ?", "nate").Find(&news)




	//e, _ := json.Marshal(news)

	//c.Header("Content-Type", "application/json")
	c.HTML(http.StatusOK, "news", news)
}
