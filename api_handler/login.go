package api_handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
)

func Login(c *gin.Context) {

	logFlag := service.GetLoginFlag(c)

	//already log user handling
	if logFlag{
		c.HTML(http.StatusOK, "home", gin.H{
			const_val.LoginFlag : service.GetLoginFlag(c),
		})
		return
	}

	c.HTML(http.StatusOK, "login", gin.H{
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
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

	c.HTML(http.StatusOK, "home",gin.H{
		const_val.LoginFlag : service.GetLoginFlag(c),
	})
}

func createSession(c *gin.Context, user *model.User) error{
	session := sessions.Default(c)
	session.Set(const_val.UserKey, user.UserId)
	err := session.Save()
	return err
}