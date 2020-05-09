package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)

func HomeGet(c *gin.Context) {
	c.HTML(http.StatusOK, const_val.TmplFileHome, gin.H{
		const_val.TmplVarLoginFlag: service.GetLoginFlag(c),
	})
}
