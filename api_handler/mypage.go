package api_handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)




func MyPageGet(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))

	reviewTemplates := service.GetReviewTemplates(userId)


	c.HTML(http.StatusOK, "myPage",gin.H{
		"userId"  : userId,
		"reviews" : reviewTemplates,
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}