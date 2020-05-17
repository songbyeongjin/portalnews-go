package impl

import (
	"github.com/stretchr/testify/assert"
	"portal_news/domain/model"
	"testing"
	"time"
)

func TestNewUserService(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(testUserS)
	assertion.IsType(new(userService), testUserS)
}

func TestUserExistCheck(t *testing.T) {
	assertion := assert.New(t)
	user := &model.User{}

	user.UserId = "test"
	assertion.True(testUserS.UserExistCheck(user))

	user.UserId = "not_exist_id" //not exist user id
	assertion.False(testUserS.UserExistCheck(user))
}

func TestCreateUser(t *testing.T) {
	assertion := assert.New(t)
	user := &model.User{}

	now := time.Now()
	user.UserId = "test" + now.String()

	assertion.False(testUserS.UserExistCheck(user))

	testUserS.CreateUser(user)

	assertion.True(testUserS.UserExistCheck(user))
}