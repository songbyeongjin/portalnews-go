package service

import "github.com/gin-gonic/gin"

type LogoutService interface {
	DeleteSession(c *gin.Context) error
}
