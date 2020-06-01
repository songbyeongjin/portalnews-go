package controller

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"log"
	"net/http"
	"portal_news/common"
	"portal_news/domain/model"
	"portal_news/service"
	"portal_news/service/impl"
)

type LoginController struct {
	LoginService service.LoginService
}

func NewLoginController(loginService service.LoginService) LoginController {
	loginController := LoginController{LoginService: loginService}
	return loginController
}

func (l LoginController)LoginGet(c *gin.Context) {
	logFlag := common.GetLoginFlag(c)

	//already log user handling
	if logFlag {
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})

		return
	}

	c.HTML(http.StatusOK, common.TmplFileLogin, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarSignUpId:  nil,
	})
}

func (l LoginController) LoginPost(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"binding error": err.Error()})
		fmt.Println(err)
		return
	}

	idNotExist, userObj := l.LoginService.UserNotExistCheck(user)
	if idNotExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorized id not exist"})
		return
	} else {
		if userObj.UserPass != user.UserPass {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized password err"})
			return
		}
	}

	if l.LoginService.CreateSession(c, user) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}

func (l LoginController) GoogleOauthGet(c *gin.Context) {
	state := l.LoginService.OauthSetCookie(c)
	url := impl.OauthGoogleConfig.AuthCodeURL(state)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (l LoginController) GoogleOauthCallbackGet(c *gin.Context) {
	cookieState, _ := c.Cookie(common.OauthGoogleCookieName)
	googleState := c.Request.URL.Query().Get(common.StateCookie)

	if cookieState != googleState {
		log.Printf("invalid google oauth state")
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})
	}

	googleCode := c.Request.URL.Query().Get("code")
	googleUser, err := l.LoginService.GetGoogleUserInfo(c, googleCode)

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
			common.TmplVarLoginFlag: common.GetLoginFlag(c),
		})
	}

	user := l.LoginService.GoogleUserDbInsert(googleUser)

	if l.LoginService.CreateSession(c, user) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}
