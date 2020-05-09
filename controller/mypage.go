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

	c.HTML(http.StatusOK, const_val.TmplFileMypage, gin.H{
		const_val.TmplVarUserId:    userId,
		const_val.TmplVarReviews:   reviewTemplates,
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
	})
}
