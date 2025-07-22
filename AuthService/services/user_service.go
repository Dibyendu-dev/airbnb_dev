package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser()	error
}

type UserServiceImpl struct {
	userRepository db.UserRepository //this depends on interface ,not userServiceImpl(class) *DI
}

func NewUserService(_userRepository db.UserRepository) UserService { //constructor fn
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("fetching user in user service")
	u.userRepository.GetById()
	return nil
}

func (u *UserServiceImpl) CreateUser() error{
	fmt.Println("creating user in userService")
	password := "hashuser2pass123"
	hashedPassword,err :=utils.HashPassword(password)
	if err != nil{
		fmt.Println("error to hashed password",err)
		return err
	}
	u.userRepository.Create(
		"user_test2","user2@test.com",hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) LoginUser() error{
	response :=utils.CheckPasswordHash("hashuser2pass123","")
	fmt.Println("Login response:",response)
	return nil
}