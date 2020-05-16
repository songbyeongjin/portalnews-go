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

	logoutS := NewLogoutService()

	assertion.NotNil(logoutS)
	assertion.IsType(new(logoutService), logoutS)
}

func TestClearSession(t *testing.T) {
	assertion := assert.New(t)
	logoutS := NewLogoutService()

	r := gin.Default()
	store := cookie.NewStore([]byte(common.SessionKey))
	r.Use(sessions.Sessions("mySession", store))


	r.GET("/clear-session", func(context *gin.Context) {
		assertion.Nil(logoutS.ClearSession(context))
	})

	res2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/clear-session", nil)
	r.ServeHTTP(res2, req2)
}
