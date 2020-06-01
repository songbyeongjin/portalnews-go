package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"portal_news/common"
	"portal_news/domain/model"
	"portal_news/service"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	userController := UserController{UserService: userService}
	return userController
}

func (u UserController) SignUpPost(c *gin.Context) {
	user := &model.User{}
	if err := c.Bind(user); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"binding error": err.Error()})
		return
	}


	isExist := u.UserService.UserExistCheck(user)
	if isExist {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id already exist"})
		return
	}

	u.UserService.CreateUser(user)

	c.HTML(http.StatusOK, common.TmplFileLogin, gin.H{
		common.TmplVarLoginFlag: common.GetLoginFlag(c),
		common.TmplVarSignUpId:  user.UserId,
	})
}