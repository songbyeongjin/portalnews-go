package impl

import (
	"portal_news/infra/db"
	"portal_news/infra/db/repository_impl"
)

var testPsDb = db.NewTestDbHandler() //use test DB

var testUserR = repository_impl.NewUserRepository(testPsDb)
var testNewsR = repository_impl.NewNewsRepository(testPsDb)
var testReviewR = repository_impl.NewReviewRepository(testPsDb)
var testRankingNewsR = repository_impl.NewRankingNewsRepository(testPsDb)


var testLoginS = NewLoginService(testUserR)
var testMyPageS = NewMyPageService(testReviewR, testNewsR)
var testLogoutS = NewLogoutService()
var testRankingNewsS = NewRankingNewsService(testRankingNewsR)
var testUserS = NewUserService(testUserR)
var testSearchS = NewSearchService(testNewsR)
var testReviewS = NewReviewService(testReviewR, testNewsR)
