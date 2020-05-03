package api_handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login", nil)
}

func LoginAuth(c *gin.Context) {
	userId := c.PostForm("userId")
	userPass := c.PostForm("userPass")

	user := &model.User{
		UserId: userId,
		UserPass: userPass,
	}

	var userObj = new(model.User)

	idNotExist := db.Instance.Where("user_id = ?", user.UserId).First(userObj).RecordNotFound()
	if idNotExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized id not exist"})
		return

	}else{
		if userObj.UserPass != user.UserPass{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized password err"})
			return
		}
	}

	if createSession(c, user) != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.HTML(http.StatusOK, "home", nil)
}

func createSession(c *gin.Context, user *model.User) error{
	session := sessions.Default(c)
	session.Set(service.UserKey, user.UserId)
	err := session.Save()
	return err
}