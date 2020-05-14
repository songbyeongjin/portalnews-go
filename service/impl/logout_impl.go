package impl

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"portal_news/service"
)

type logoutService struct {
}

func NewLogoutService() service.LogoutService {
	logoutService := logoutService{}
	return &logoutService
}

func (l *logoutService) DeleteSession(c *gin.Context) error{
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	return err
}

