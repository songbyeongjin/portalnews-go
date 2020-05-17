package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMyPageService(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(testMyPageS)
	assertion.IsType(new(myPageService), testMyPageS)
}

func TestGetReviewByUserId(t *testing.T) {
	assertion := assert.New(t)

	existUserId := "test" //already exist "test" user and "test" user's review
	assertion.NotEmpty(*testMyPageS.GetReviewByUserId(existUserId))

	notExistUserId := "not_exist_id" //not exist user id
	assertion.Empty(*testMyPageS.GetReviewByUserId(notExistUserId))
}