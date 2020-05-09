package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
)

func NaverGet(c *gin.Context) {
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", const_val.Naver).Find(&rankingNews)

	c.HTML(http.StatusOK, const_val.TmplFileNews, gin.H{
		const_val.TmplVarNews:      rankingNews,
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		const_val.TmplVarLanguage:  const_val.TmplVarJapanese,
	})
}

func DaumGet(c *gin.Context) {
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", const_val.Daum).Find(&rankingNews)

	c.HTML(http.StatusOK, const_val.TmplFileNews, gin.H{
		const_val.TmplVarNews:      rankingNews,
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		const_val.TmplVarLanguage:  const_val.TmplVarKorean,
	})
}

func NateGet(c *gin.Context) {
	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", const_val.Nate).Find(&rankingNews)

	c.HTML(http.StatusOK, const_val.TmplFileNews, gin.H{
		const_val.TmplVarNews:      rankingNews,
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		const_val.TmplVarLanguage:  const_val.TmplVarKorean,
	})
}

func NaverLanguageGet(c *gin.Context) {
	language := c.Param("language")

	rankingNews := []model.RankingNews{}

	db.Instance.Where("portal = ?", const_val.Naver).Find(&rankingNews)

	c.HTML(http.StatusOK, const_val.TmplFileNews, gin.H{
		const_val.TmplVarNews:      rankingNews,
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		const_val.TmplVarLanguage:  language,
	})
}
