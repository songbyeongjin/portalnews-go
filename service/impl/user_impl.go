package impl

import (
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/service"
)

type userService struct {
	userRepository   repository_interface.UserRepository
}

func NewUserService(userRepository repository_interface.UserRepository) service.UserService {
	userService := userService{userRepository: userRepository}

	return &userService
}

func (u *userService) UserExistCheck(user *model.User) bool{
	userRet := u.userRepository.FindFirstByUserId(user.UserId)

	if userRet != nil{
		return true
	}else{
		return false
	}
}

func (u *userService) CreateUser(user *model.User){
	u.userRepository.Create(user)
}