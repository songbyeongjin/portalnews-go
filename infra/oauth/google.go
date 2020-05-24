package oauth

import (
	"fmt"
	"golang.org/x/oauth2/google"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"portal_news/service/impl"
)

type GoogleApiInfo struct {
	ID   string `yaml:"oauth_google_id"`
	Pass string `yaml:"oauth_google_pass"`
}

func SetOauthGoogleConfig() error{
	OauthGoogleInfo, err := GetOauthGoogleInfo()
	if err != nil{
		fmt.Println(err)
		return err
	}

	impl.OauthGoogleConfig.RedirectURL = "http://korea-portal-news.com/login/google-oauth/callback"
	impl.OauthGoogleConfig.ClientID = OauthGoogleInfo.ID
	impl.OauthGoogleConfig.ClientSecret = OauthGoogleInfo.Pass
	impl.OauthGoogleConfig.Scopes = []string{"https://www.googleapis.com/auth/userinfo.email"}
	impl.OauthGoogleConfig.Endpoint = google.Endpoint
	return nil
}

func GetOauthGoogleInfo() (*GoogleApiInfo, error) {
	//temporary path for debug mode
	buf, err := ioutil.ReadFile(`./oauth_google.yaml`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var oauthGoogleInfo *GoogleApiInfo

	err = yaml.Unmarshal(buf, &oauthGoogleInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return oauthGoogleInfo, nil
}
