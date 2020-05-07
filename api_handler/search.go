package api_handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)




func SearchGet(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))

	reviewTemplates := service.GetReviewTemplates(userId)


	c.HTML(http.StatusOK, "search",gin.H{
		"userId"  : userId,
		"reviews" : reviewTemplates,
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}

func SearchNewsGet(c *gin.Context) {
	session := sessions.Default(c)
	urlValue := c.Request.URL.Query()

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	news, lang := service.GetSearchNews(urlValue)


	c.HTML(http.StatusOK, "search",gin.H{
		"userId"  : userId,
		const_val.LoginFlag : service.GetLoginFlag(c),
		const_val.News : news,
		"language" : lang,
	})
}