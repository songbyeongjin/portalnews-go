package api_handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"portal_news/const_val"
	"portal_news/model"
	"portal_news/service"
)

func LoginGet(c *gin.Context) {

	logFlag := service.GetLoginFlag(c)

	//already log user handling
	if logFlag {
		c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
			const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		})
		return
	}

	c.HTML(http.StatusOK, const_val.TmplFileLogin, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		const_val.TmplVarSignUpId:  nil,
	})
}

func LoginAuthPost(c *gin.Context) {
	userId := c.PostForm("userId")
	userPass := c.PostForm("userPass")

	user := &model.User{
		UserId:   userId,
		UserPass: userPass,
	}

	//validation check
	if !service.UserValidation(user) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "parameter error"})
		return
	}

	idNotExist, userObj := service.UserExistCheck(user)
	if idNotExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized id not exist"})
		return
	} else {
		if userObj.UserPass != user.UserPass {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized password err"})
			return
		}
	}

	if service.CreateSession(c, user) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
	})
}

func GoogleOauthGet(c *gin.Context) {
	state := service.OauthSetCookie(c)
	url := service.OauthGoogleConfig.AuthCodeURL(state)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func GoogleOauthCallbackGet(c *gin.Context) {
	cookieState, _ := c.Cookie(const_val.OauthGoogleCookieName)
	googleState := c.Request.URL.Query().Get(const_val.StateCookie)

	if cookieState != googleState {
		log.Printf("invalid google oauth state")
		c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
			const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		})
	}

	googleCode := c.Request.URL.Query().Get("code")
	googleUser, err := service.GetGoogleUserInfo(c, googleCode)

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
			const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		})
	}

	user := service.GoogleUserDbInsert(googleUser)

	if service.CreateSession(c, user) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
	})
}
