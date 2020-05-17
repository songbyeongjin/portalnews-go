package impl

import (
	"github.com/stretchr/testify/assert"
	"portal_news/common"
	"testing"
)

func TestNewRankingNewsService(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(testRankingNewsS)
	assertion.IsType(new(rankingNewsService), testRankingNewsS)
}

func TestGetNewsByPortal(t *testing.T) {
	assertion := assert.New(t)

	portal := common.Naver
	assertion.NotEmpty(testRankingNewsS.GetNewsByPortal(portal))

	portal = common.Daum
	assertion.NotEmpty(testRankingNewsS.GetNewsByPortal(portal))

	portal = common.Nate
	assertion.NotEmpty(testRankingNewsS.GetNewsByPortal(portal))
}