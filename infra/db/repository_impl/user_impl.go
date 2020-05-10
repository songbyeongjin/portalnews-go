package repository_impl

import (
	"portal_news/domain/model"
	"portal_news/domain/repository_interface"
	"portal_news/infra/db"
)

type UserRepository struct {
	dbHandler db.Handler
}

func NewUserRepository(dbHandler db.Handler) repository_interface.UserRepository {
	userRepository := UserRepository{dbHandler}
	return &userRepository
}


func (n *UserRepository) FindFirstByUserId(userId string)  *model.User {
	user := &model.User{}
	n.dbHandler.Conn.Where("user_id = ?", userId).Find(user)

	if user.UserId == ""{
		return nil
	}

	return user
}

func (n *UserRepository) Create(user *model.User) {
	n.dbHandler.Conn.Create(user)
}