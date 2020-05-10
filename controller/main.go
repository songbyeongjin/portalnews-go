package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
)

type MainController struct {
}

func NewMainController() MainController {
	mainController := MainController{}

	return mainController
}

func (m MainController) HomeGet(c *gin.Context) {
	c.HTML(http.StatusOK, common.TmplFileHome, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
	})
}