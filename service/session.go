package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(const_val.UserKey)
		if user == nil {
			// Abort the request with the appropriate error code
			c.HTML(http.StatusUnauthorized, "notLogin",gin.H{
				const_val.LoginFlag : GetLoginFlag(c),
			})

			c.Abort()
			return
		}else{
			c.Set(const_val.UserKey, user) // ユーザidをセット
			c.Next()
		}
	}
}
