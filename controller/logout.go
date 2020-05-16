package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
	"portal_news/service"
)

type LogoutController struct {
	LogoutService service.LogoutService
}

func NewLogoutController(logoutService service.LogoutService) LogoutController {
	loginController := LogoutController{LogoutService: logoutService}
	return loginController
}

func (l LogoutController) LogoutGet(c *gin.Context) {
	if l.LogoutService.ClearSession(c) != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session delete failed"})
		return
	}

	c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}