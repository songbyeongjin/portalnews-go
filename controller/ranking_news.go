package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
	"portal_news/service"
)

type RankingNewsController struct {
	RankingNewsService service.RankingNewsService
}

func NewRankingNewsController(rankingNewsService service.RankingNewsService) RankingNewsController {
	rankingNewsController := RankingNewsController{RankingNewsService: rankingNewsService}
	return rankingNewsController
}

func (r RankingNewsController) NaverGet(c *gin.Context) {
	naverNews := r.RankingNewsService.GetNewsByPortal(common.Naver)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      naverNews,
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarLanguage:  common.TmplVarJapanese,
	})
}

func (r RankingNewsController) DaumGet(c *gin.Context) {
	daumNews := r.RankingNewsService.GetNewsByPortal(common.Daum)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      daumNews,
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarLanguage:  common.TmplVarKorean,
	})
}

func (r RankingNewsController) NateGet(c *gin.Context) {
	nateNews := r.RankingNewsService.GetNewsByPortal(common.Nate)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      nateNews,
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarLanguage:  common.TmplVarKorean,
	})
}

func (r RankingNewsController) NaverLanguageGet(c *gin.Context) {
	language := c.Param("language")

	naverNew := r.RankingNewsService.GetNewsByPortal(common.Naver)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      naverNew,
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarLanguage:  language,
	})
}
