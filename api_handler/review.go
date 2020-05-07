package api_handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
	"strings"
	"time"
)

type Jsonody struct{
	ReviewTitle string `json:"reviewTitle"`
	ReviewContent string `json:"reviewContent"`
	NewsUrl string `json:"newsUrl"`
}

func WriteReviewGET(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := const_val.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+ len(targetStr):]


	review, modifyFlag := service.GetCreateReviewTemplate(userId, newsUrl)

	if review == nil {
		c.HTML(http.StatusOK, "home",gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}


	c.HTML(http.StatusOK, "writeReview",gin.H{
		"userId": userId,
		"review" : review,
		"modifyFlag" : modifyFlag,
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}

func WriteReviewPOST(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	bodyJson := &Jsonody{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, bodyJson)

	if err != nil{
		c.HTML(http.StatusOK, "home",gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}
	
	review := &model.Review{}
	notExist := db.Instance.Where("news_url = ? AND user_id = ?", bodyJson.NewsUrl, userId).First(review).RecordNotFound()
	if !notExist {
		db.Instance.Model(review).Updates(map[string]interface{}{"title":bodyJson.ReviewTitle, "content":bodyJson.ReviewContent})
	}else{
		createReview := &model.Review{
			UserId : userId,
			NewsUrl: bodyJson.NewsUrl,
			Title:   bodyJson.ReviewTitle,
			Content: bodyJson.ReviewContent,
			Date:    time.Now(),
		}
		db.Instance.Create(createReview)
	}

	c.JSON(http.StatusOK, "/mypage")
}



func WriteReviewPUT(c *gin.Context) {
	session := sessions.Default(c)

	bodyJson := &Jsonody{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(body))
	err := json.Unmarshal(body, bodyJson)

	if err != nil{
		c.HTML(http.StatusOK, "home",gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := const_val.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+ len(targetStr):]

	var review = &model.Review{}
	notExist :=  db.Instance.Where("news_url = ? AND user_id = ?", newsUrl, userId).First(review).RecordNotFound()

	if notExist {
		c.HTML(http.StatusOK, "home",gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}else{
		db.Instance.Model(review).Updates(map[string]interface{}{"title":bodyJson.ReviewTitle, "content":bodyJson.ReviewContent})
	}

	c.JSON(http.StatusOK, "/mypage")
}

func WriteReviewDELETE(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := const_val.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+ len(targetStr):]

	var review = &model.Review{}
	notExist :=  db.Instance.Where("news_url = ? AND user_id = ?", newsUrl, userId).First(review).RecordNotFound()

	if notExist {
		c.HTML(http.StatusOK, "home",gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}else{
		db.Instance.Delete(review)
	}

	c.JSON(http.StatusOK, "/mypage")
}