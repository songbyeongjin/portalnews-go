package impl

import (
	"github.com/stretchr/testify/assert"
	"portal_news/domain/model"
	"testing"
)

func TestNewLoginService(t *testing.T) {
	assertion := assert.New(t)

	loginServiceS := NewLoginService(nil)

	assertion.NotNil(loginServiceS)
	assertion.IsType(new(loginService), loginServiceS)
}

func TestUserValidation(t *testing.T){
	assertion := assert.New(t)

	loginServiceS := NewLoginService(nil)

	assertion.False(loginServiceS.UserValidation(nil))

	user := &model.User{
	}

	user.UserId = "testId"
	user.UserPass = "testPass"
	user.Oauth = "testOauth"
	assertion.True(loginServiceS.UserValidation(user))

	user.UserId = ""
	user.UserPass = "testPass"
	user.Oauth = "testOauth"
	assertion.False(loginServiceS.UserValidation(user))

	user.UserId = "testId"
	user.UserPass = ""
	user.Oauth = "testOauth"
	assertion.False(loginServiceS.UserValidation(user))

	user.UserId = "testId"
	user.UserPass = "testPass"
	user.Oauth = ""
	assertion.True(loginServiceS.UserValidation(user))
}