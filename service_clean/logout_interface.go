package service_clean

import "github.com/gin-gonic/gin"

type LogoutService interface {
	DeleteSession(c *gin.Context) error
}
