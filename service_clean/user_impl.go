package service_clean

import (
	"portal_news/domain_clean/model"
	"portal_news/domain_clean/repository_interface"
)

type userService struct {
	userRepository   repository_interface.UserRepository
}

func NewUserService(userRepository repository_interface.UserRepository) UserService {
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