package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReviewService(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(testReviewS)
	assertion.IsType(new(reviewService), testReviewS)
}

func TestGetReviewByNewsUrlAndUserId(t *testing.T) {
	assertion := assert.New(t)

	existUrl := "test1"
	existUserId := "test"

	notExistUrl := "not_exist_url"
	notExistUserId := "not_exist_user_id"

	r, m := testReviewS.GetReviewByNewsUrlAndUserId(notExistUserId, notExistUrl)
	assertion.Nil(r)
	assertion.False(m)

	r, m = testReviewS.GetReviewByNewsUrlAndUserId(existUserId, existUrl)
	assertion.NotNil(r)
	assertion.True(m)

	r, m = testReviewS.GetReviewByNewsUrlAndUserId(notExistUserId, existUrl)
	assertion.NotNil(r)
	assertion.False(m)
}