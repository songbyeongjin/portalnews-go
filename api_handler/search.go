package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)

func SearchGet(c *gin.Context) {
	c.HTML(http.StatusOK, const_val.TmplFileSerch, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
	})
}

func SearchNewsGet(c *gin.Context) {
	urlValue := c.Request.URL.Query()
	news, lang := service.GetSearchNews(urlValue)

	c.HTML(http.StatusOK, const_val.TmplFileSerch, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
		const_val.TmplVarNews:      news,
		const_val.TmplVarLanguage:  lang,
	})
}
