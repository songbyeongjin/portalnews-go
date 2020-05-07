package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)

func HomeGet(c *gin.Context){
	c.HTML(http.StatusOK, const_val.TemplateHome,gin.H{
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}
