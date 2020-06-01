package impl

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"portal_news/common"
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/service"
)
var OauthGoogleConfig oauth2.Config

type loginService struct {
	userRepository repository_interface.UserRepository
}

func NewLoginService(userRepository repository_interface.UserRepository) service.LoginService {
	loginService := loginService{userRepository: userRepository}

	return &loginService
}


func (l *loginService) UserNotExistCheck(user *model.User) (bool, *model.User){
	userRet := l.userRepository.FindFirstByUserId(user.UserId)

	if userRet == nil{
		return true, userRet
	}else{
		return false, userRet
	}
}

func (l *loginService) CreateSession(c *gin.Context, user *model.User) error{
	session := sessions.Default(c)
	session.Set(common.UserKey, user.UserId)
	err := session.Save()
	return err
}

func (l *loginService) OauthSetCookie(c *gin.Context) string{
	randomByte := make([]byte, 16)
	rand.Read(randomByte)
	state := base64.URLEncoding.EncodeToString(randomByte)

	c.SetCookie(common.OauthGoogleCookieName, state, 60*24, "", "", false, false)
	return state
}

func (l *loginService) GetGoogleUserInfo(c *gin.Context, code string) (*service.OauthGoogleUser, error){
	token, tokenErr := OauthGoogleConfig.Exchange(c, code)

	if tokenErr != nil {
		return nil, fmt.Errorf("failed to Exchange %s", tokenErr.Error())
	}

	resp, UserErr := http.Get(common.OauthGoogleUrlAPI + token.AccessToken)

	if UserErr != nil {
		return nil, fmt.Errorf("failed to Get User Info %s", UserErr.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	googleUser := &service.OauthGoogleUser{}

	jsonErr := json.Unmarshal(body, googleUser)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return googleUser, nil
}

func (l *loginService) GoogleUserDbInsert(googleUser *service.OauthGoogleUser)*model.User{
	user := &model.User{
		UserId: googleUser.Email,
		Oauth:  "google",
	}

	idNotExist, _ := l.UserNotExistCheck(user)
	if idNotExist {
		//insert user
		l.userRepository.Create(user)
	}

	return user
}




