package api_handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
	"strings"
	"time"
)

func WriteReviewGET(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := const_val.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+ len(targetStr):]

	var news = &model.News{}
	notExist :=  db.Instance.Where("url = ?", newsUrl).First(news).RecordNotFound()

	if notExist {
		c.HTML(http.StatusOK, "home",gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}

	c.HTML(http.StatusOK, "writeReview",gin.H{
		"userId": userId,
		"newsPortal": news.Portal,
		"newsPress": news.Press,
		"newsTitle": news.Title,
		const_val.LoginFlag : service.GetLoginFlag(c),
		"newsUrl":news.Url,
	})
}

func WriteReviewPOST(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	title := c.PostForm("title")
	content := c.PostForm("content")
	url := c.PostForm("url")

	review := &model.Review{
		UserId : userId,
		NewsUrl: url,
		Title:   title,
		Content: content,
		Date:    time.Now(),
	}

	db.Instance.Create(review)
	c.Redirect(http.StatusMovedPermanently, "/mypage")
}