package service_clean

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type logoutService struct {
}

func NewLogoutService() LogoutService {
	logoutService := logoutService{}
	return &logoutService
}

func (l *logoutService) DeleteSession(c *gin.Context) error{
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	return err
}
