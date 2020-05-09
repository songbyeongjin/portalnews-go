package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)

func LogoutGet(c *gin.Context) {
	if service.DeleteSession(c) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session delete failed"})
		return
	}

	c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
	})
}
