package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser() error
}

type UserServiceImpl struct {
	userRepository db.UserRepository //this depends on interface ,not userServiceImpl(class) *DI
}

func NewUserService(_userRepository db.UserRepository) UserService { //constructor fn
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("creating user in user service")
	u.userRepository.Create()
	return nil
}
