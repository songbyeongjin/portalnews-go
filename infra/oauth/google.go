package oauth

import (
	"fmt"
	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"portal_news/service"
)

type OauthGoogleApiInfo struct {
	ID   string `yaml:"oauth_google_id"`
	Pass string `yaml:"oauth_google_pass"`
}

func SetOauthGoogleConfig() {
	OauthGoogleInfo, _ := GetOauthGoogleInfo()
	service.OauthGoogleConfig.RedirectURL = "http://localhost:8080/login/google-oauth/callback"
	service.OauthGoogleConfig.ClientID = OauthGoogleInfo.ID
	service.OauthGoogleConfig.ClientSecret = OauthGoogleInfo.Pass
	service.OauthGoogleConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	service.OauthGoogleConfig.Endpoint = google.Endpoint
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
