package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
	"portal_news/service"
)

type SearchController struct {
	SearchService service.SearchService
}

func NewSearchController(searchService service.SearchService) SearchController {
	searchController := SearchController{SearchService: searchService}
	return searchController
}

func (s SearchController) SearchGet(c *gin.Context) {
	c.HTML(http.StatusOK, common.TmplFileSerch, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}

func (s SearchController)SearchNewsGet(c *gin.Context) {
	urlValue := c.Request.URL.Query()
	news, lang := s.SearchService.GetSearchNews(urlValue)

	c.HTML(http.StatusOK, common.TmplFileSerch, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarNews:      news,
		common.TmplVarLanguage:  lang,
	})
}
