package controller

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
	"portal_news/service"
)

type MyPageController struct {
	MyPageService service.MyPageService
}

func NewMyPageController(myPageService service.MyPageService) MyPageController {
	myPageController := MyPageController{MyPageService: myPageService}
	return myPageController
}

func (m MyPageController) MyPageGet(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(common.UserKey))

	reviewTemplates := m.MyPageService.GetReviewByUserId(userId)

	c.HTML(http.StatusOK, common.TmplFileMypage, gin.H{
		common.TmplVarUserId:    userId,
		common.TmplVarReviews:   reviewTemplates,
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}