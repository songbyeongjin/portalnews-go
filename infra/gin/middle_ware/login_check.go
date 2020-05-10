package middleWare

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(common.UserKey)
		if user == nil {
			// Abort the request with the appropriate error code
			c.HTML(http.StatusUnauthorized, common.TmplFileNotLogin, gin.H{
				common.TmplVarLoginFlag: common.GetLoginFlag(c),
			})

			c.Abort()
			return
		} else {
			c.Set(common.UserKey, user)
			c.Next()
		}
	}
}