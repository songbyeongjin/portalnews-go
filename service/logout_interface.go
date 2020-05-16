package service

import "github.com/gin-gonic/gin"

type LogoutService interface {
	ClearSession(c *gin.Context) error
}
