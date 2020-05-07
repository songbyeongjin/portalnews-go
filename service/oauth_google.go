package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
)

var OauthGoogleConfig oauth2.Config

type OauthGoogleApiInfo struct {
	ID   string `yaml:"oauth_google_id"`
	Pass string `yaml:"oauth_google_pass"`
}

type OauthGoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func SetOauthGoogleConfig() {
	OauthGoogleInfo, _ := GetOauthGoogleInfo()
	OauthGoogleConfig.RedirectURL = "http://localhost:8080/google-oauth/callback"
	OauthGoogleConfig.ClientID = OauthGoogleInfo.ID
	OauthGoogleConfig.ClientSecret = OauthGoogleInfo.Pass
	OauthGoogleConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	OauthGoogleConfig.Endpoint = google.Endpoint
}

func OauthSetCookie(c *gin.Context) string {
	randomByte := make([]byte, 16)
	rand.Read(randomByte)
	state := base64.URLEncoding.EncodeToString(randomByte)

	c.SetCookie(const_val.OauthGoogleCookieName, state, 60*24, "", "", false, false)
	return state
}

func GetGoogleUserInfo(c *gin.Context, code string) (*OauthGoogleUser, error) {
	token, tokenErr := OauthGoogleConfig.Exchange(c, code)

	if tokenErr != nil {
		return nil, fmt.Errorf("failed to Exchange %s", tokenErr.Error())
	}

	resp, UserErr := http.Get(const_val.OauthGoogleUrlAPI + token.AccessToken)

	if UserErr != nil {
		return nil, fmt.Errorf("failed to Get User Info %s", UserErr.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	googleUser := &OauthGoogleUser{}

	jsonErr := json.Unmarshal(body, googleUser)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return googleUser, nil
}

func GetOauthGoogleInfo() (*OauthGoogleApiInfo, error) {
	//temporary path for debug mode
	buf, err := ioutil.ReadFile(`C:\Users\SONG\Documents\study\go\src\portal_news\oauth_google.yaml`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var oauthGoogleInfo *OauthGoogleApiInfo

	err = yaml.Unmarshal(buf, &oauthGoogleInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return oauthGoogleInfo, nil
}

func GoogleUserDbInsert(googleUser *OauthGoogleUser) *model.User {
	user := &model.User{
		UserId: googleUser.Email,
		Oauth:  "google",
	}

	idNotExist, _ := UserExistCheck(user)
	if idNotExist {
		//insert user
		db.Instance.Create(user)
	}

	return user
}
