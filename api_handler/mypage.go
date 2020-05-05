package api_handler

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
)

func MyPageGet(c *gin.Context) {
	session := sessions.Default(c)

	userId := fmt.Sprintf("%v", session.Get(const_val.UserKey))

	c.JSON(http.StatusOK, gin.H{"state": "welcome " + userId})
}