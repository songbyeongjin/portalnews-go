package api_handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)

func Logout(c *gin.Context) {
	if deleteSession(c) != nil{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session delete failed"})
		return
	}

	c.HTML(http.StatusOK, "home", gin.H{
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}

func deleteSession(c *gin.Context) error{
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	return err
}