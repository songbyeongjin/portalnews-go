package impl

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"portal_news/common"
	"testing"
)



func TestNewLogoutService(t *testing.T) {
	assertion := assert.New(t)


	assertion.NotNil(testLogoutS)
	assertion.IsType(new(logoutService), testLogoutS)
}

func TestClearSession(t *testing.T) {
	assertion := assert.New(t)

	r := gin.Default()
	store := cookie.NewStore([]byte(common.SessionKey))
	r.Use(sessions.Sessions("mySession", store))


	r.GET("/clear-session", func(context *gin.Context) {
		assertion.Nil(testLogoutS.ClearSession(context))
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/clear-session", nil)
	r.ServeHTTP(res, req)
}
