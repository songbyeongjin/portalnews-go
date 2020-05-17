package impl

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"portal_news/common"
	"portal_news/domain/model"
	"portal_news/service"
	"testing"
	"time"
)

func TestNewLoginService(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(testLoginS)
	assertion.IsType(new(loginService), testLoginS)
}

func TestUserValidation(t *testing.T){
	assertion := assert.New(t)

	assertion.False(testLoginS.UserValidation(nil))

	user := &model.User{
	}

	user.UserId = "testId"
	user.UserPass = "testPass"
	user.Oauth = "testOauth"
	assertion.True(testLoginS.UserValidation(user))

	user.UserId = ""
	user.UserPass = "testPass"
	user.Oauth = "testOauth"
	assertion.False(testLoginS.UserValidation(user))

	user.UserId = "testId"
	user.UserPass = ""
	user.Oauth = "testOauth"
	assertion.False(testLoginS.UserValidation(user))

	user.UserId = "testId"
	user.UserPass = "testPass"
	user.Oauth = ""
	assertion.True(testLoginS.UserValidation(user))
}

func TestUserNotExistCheck(t *testing.T){
	assertion := assert.New(t)

	existUser := &model.User{
		UserId: "test", // already exist "test" user in test db
	}

	assertion.False(testLoginS.UserNotExistCheck(existUser))

	notExistUser := &model.User{
		UserId: "not_exist_id", // already exist "test" user in test db
	}

	assertion.True(testLoginS.UserNotExistCheck(notExistUser))
}

func TestCreateSession(t *testing.T){
	assertion := assert.New(t)

	r := gin.Default()
	store := cookie.NewStore([]byte(common.SessionKey))
	r.Use(sessions.Sessions("mySession", store))

	user := &model.User{
		UserId: "test",
	}

	r.GET("/create-session", func(context *gin.Context) {
		assertion.False(common.GetLoginFlag(context))
		assertion.Nil(testLoginS.CreateSession(context, user))
		assertion.True(common.GetLoginFlag(context))
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/create-session", nil)
	r.ServeHTTP(res, req)

}

func TestOauthSetCookie(t *testing.T) {
	assertion := assert.New(t)

	r := gin.Default()
	store := cookie.NewStore([]byte(common.SessionKey))
	r.Use(sessions.Sessions("mySession", store))
	r.GET("/oauth-set-cookie", func(context *gin.Context) {
		setCookieMap := context.Writer.Header()["Set-Cookie"]
		assertion.Nil(setCookieMap)

		state := testLoginS.OauthSetCookie(context)
		assertion.NotNil(state)

		setCookieMap = context.Writer.Header()["Set-Cookie"]
		assertion.NotNil(setCookieMap)
	})

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/oauth-set-cookie", nil)
	r.ServeHTTP(res, req)
}

func TestGetGoogleUserInfo(t *testing.T){
	assertion := assert.New(t)
	now := time.Now()

	userId := "test" + now.String()

	user := &model.User{
		UserId: userId,
	}

	googleUser := &service.OauthGoogleUser{}
	googleUser.Email = userId

	userNotExist, _ := testLoginS.UserNotExistCheck(user)
	assertion.True(userNotExist)

	testLoginS.GoogleUserDbInsert(googleUser)

	userNotExist, _ = testLoginS.UserNotExistCheck(user)
	assertion.False(userNotExist)
}