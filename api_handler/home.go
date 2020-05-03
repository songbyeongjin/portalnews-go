package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/service"
)

func Home(c *gin.Context){
	c.HTML(http.StatusOK, "home",gin.H{
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}
