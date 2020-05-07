package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/model"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(const_val.UserKey)
		if user == nil {
			// Abort the request with the appropriate error code
			c.HTML(http.StatusUnauthorized, const_val.TmplFileNotLogin, gin.H{
				const_val.TmplVarLoginFlag: GetLoginFlag(c),
			})

			c.Abort()
			return
		} else {
			c.Set(const_val.UserKey, user)
			c.Next()
		}
	}
}

func CreateSession(c *gin.Context, user *model.User) error {
	session := sessions.Default(c)
	session.Set(const_val.UserKey, user.UserId)
	err := session.Save()
	return err
}

func DeleteSession(c *gin.Context) error {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	return err
}
