package service_clean

import (
	"github.com/gin-gonic/gin"
	"portal_news/domain_clean/model"
)

type OauthGoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

type LoginService interface {
	UserValidation(user *model.User) bool
	UserNotExistCheck(user *model.User) (bool, *model.User)
	CreateSession(c *gin.Context, user *model.User) error
	OauthSetCookie(c *gin.Context) string
	GetGoogleUserInfo(c *gin.Context, code string) (*OauthGoogleUser, error)
	GoogleUserDbInsert(googleUser *OauthGoogleUser) *model.User
}
