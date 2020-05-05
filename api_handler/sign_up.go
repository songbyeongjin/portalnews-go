package api_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/const_val"
	"portal_news/db"
	"portal_news/model"
	"portal_news/service"
)

func SignUpPost(c *gin.Context) {
	userId := c.PostForm("userId")
	userPass := c.PostForm("userPass")

	user := &model.User{
		UserId: userId,
		UserPass: userPass,
	}


	idNotExist, _ := service.UserExistCheck(user)
	if !idNotExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id already exist"})
		return
	}

	db.Instance.Create(user)
	c.HTML(http.StatusOK, "login", gin.H{
		const_val.LoginFlag : service.GetLoginFlag(c),
		"signUpId" : user.UserId,
	})
}