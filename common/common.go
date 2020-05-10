package common

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	// *** Portal Name
	Naver = "naver"
	Nate  = "nate"
	Daum  = "daum"
	// Portal Name ***

	// *** Template Name
	TmplFileHome        = "home"
	TmplFileLogin       = "login"
	TmplFileNotLogin    = "notLogin"
	TmplFileNews        = "news"
	TmplFileSerch       = "search"
	TmplFileMypage      = "myPage"
	TmplFileWriteReview = "writeReview"
	// Template Name ***

	// *** Template Var
	TmplVarLoginFlag = "loginFlag"
	TmplVarNews      = "news"
	TmplVarLanguage  = "language"
	TmplVarKorean    = "korean"
	TmplVarJapanese  = "japanese"
	TmplVarSignUpId  = "signUpId"
	TmplVarUserId    = "userId"
	TmplVarReviews   = "reviews"
	TmplVarReview   = "review"
	TmplVarModiyFlag   = "modifyFlag"
	// Template Var ***


	// *** Http Url
	HttpsUrl    = `https://`
	HttpUrl     = `http://`
	// Http Url ***

	// *** Session/Oauth
	SessionKey = "secret"
	UserKey     = "user"
	StateCookie = "state"
	OauthGoogleCookieName = "googleState"
	OauthGoogleUrlAPI     = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	// Session/Oauth ***


	ReviewDeleteTargetStr = `/review/`
)

func AddHttpsString(url string) string {
	if strings.Index(url, "v.media.daum.net/") == -1 {
		return HttpsUrl + url
	} else {
		return HttpUrl + url
	}
}

func GetLoginFlag(c *gin.Context) bool {
	session := sessions.Default(c)
	user := session.Get(UserKey)
	if user != nil {
		return true
	} else {
		return false
	}
}
