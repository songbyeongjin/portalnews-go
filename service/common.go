package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"portal_news/const_val"
	"strings"
)

func AddHttpsString(url string) string {
	if strings.Index(url, "v.media.daum.net/") == -1 {
		return const_val.HttpsUrl + url
	} else {
		return const_val.HttpUrl + url
	}
}
func GetLoginFlag(c *gin.Context) bool {
	session := sessions.Default(c)
	user := session.Get(const_val.UserKey)
	if user != nil {
		return true
	} else {
		return false
	}
}
