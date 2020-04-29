package api_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/db"
	"portal_news/model"
)

func Home(c *gin.Context){
	/*
		nePortal:"song",
			wsObj := model.News{
			Id: 500,
			Title:"titleTest",
			Content:"contentTest",
			Press:"press",
			Writer:"writerTest",
			Date:time.Now(),
			Url:"wwwa.test.com",
			CreatedAt:time.Now(),
			UpdatedAt:time.Now(),
		}*/

	news := []model.News{}

	db.Instance.Find(&news)
	fmt.Println(news)

	//c.Header("Content-Type", "application/json")
	c.HTML(http.StatusOK, "home",&news)
}
