package controller_clean

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
	"portal_news/service_clean"
)

type RankingNewsController struct {
	RankingNewsService service_clean.RankingNewsService
}

func NewRankingNewsController(rankingNewsService service_clean.RankingNewsService) RankingNewsController {
	rankingNewsController := RankingNewsController{RankingNewsService: rankingNewsService}
	return rankingNewsController
}

func (r RankingNewsController) NaverGet(c *gin.Context) {
	naverNews := r.RankingNewsService.GetNewsByPortal(common.Naver)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      naverNews,
		common.TmplVarLoginFlag: service.GetLoginFlag(c),
		common.TmplVarLanguage:  common.TmplVarJapanese,
	})
}

func (r RankingNewsController) DaumGet(c *gin.Context) {
	daumNews := r.RankingNewsService.GetNewsByPortal(common.Daum)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      daumNews,
		common.TmplVarLoginFlag: service.GetLoginFlag(c),
		common.TmplVarLanguage:  common.TmplVarKorean,
	})
}

func (r RankingNewsController) NateGet(c *gin.Context) {
	nateNews := r.RankingNewsService.GetNewsByPortal(common.Nate)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      nateNews,
		common.TmplVarLoginFlag: service.GetLoginFlag(c),
		common.TmplVarLanguage:  common.TmplVarKorean,
	})
}

func (r RankingNewsController) NaverLanguageGet(c *gin.Context) {
	language := c.Param("language")

	naverNew := r.RankingNewsService.GetNewsByPortal(common.Naver)


	c.HTML(http.StatusOK, common.TmplFileNews, gin.H{
		common.TmplVarNews:      naverNew,
		common.TmplVarLoginFlag: service.GetLoginFlag(c),
		common.TmplVarLanguage:  language,
	})
}
