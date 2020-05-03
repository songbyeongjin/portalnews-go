package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MyPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "session test success"})
}