package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"portal_news/common"
	"portal_news/service"
	"strings"
)


type ReviewController struct {
	ReviewService service.ReviewService
}

func NewReviewController(reviewService service.ReviewService) ReviewController {
	newsController := ReviewController{ReviewService: reviewService}
	return newsController
}

func (r ReviewController) WriteReviewGET(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(common.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := common.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+len(targetStr):]

	review, modifyFlag := r.ReviewService.GetReviewByNewsUrlAndUserId(userId, newsUrl)

	if review == nil {
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})
		return
	}

	c.HTML(http.StatusOK, common.TmplFileWriteReview, gin.H{
		common.TmplVarUserId:    userId,
		common.TmplVarReview:    review,
		common.TmplVarModiyFlag: modifyFlag,
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}

func (r ReviewController) WriteReviewPOST(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(common.UserKey))
	jsonBody := &service.JsonBody{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, jsonBody)

	if err != nil {
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})
		return
	}

	r.ReviewService.PostReview(jsonBody, userId)

	c.JSON(http.StatusOK, "/mypage")
}

func (r ReviewController) UpdateReviewPUT(c *gin.Context) {
	session := sessions.Default(c)

	jsonBody := &service.JsonBody{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, jsonBody)

	if err != nil {
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})
		return
	}

	userId := fmt.Sprintf("%v", session.Get(common.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := common.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+len(targetStr):]

	err = r.ReviewService.UpdateReview(newsUrl, userId, jsonBody)

	if err != nil {
		log.Print(err)
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})
		return
	}

	c.JSON(http.StatusOK, "/mypage")
}

func (r ReviewController) DeleteReviewDELETE(c *gin.Context) {
	session := sessions.Default(c)
	userId := fmt.Sprintf("%v", session.Get(common.UserKey))
	queryUrl := c.Request.URL.String()
	targetStr := common.ReviewDeleteTargetStr
	targetIndex := strings.Index(queryUrl, targetStr)
	newsUrl := queryUrl[targetIndex+len(targetStr):]

	r.ReviewService.DeleteReview(newsUrl , userId)

	c.JSON(http.StatusOK, "/mypage")
}
