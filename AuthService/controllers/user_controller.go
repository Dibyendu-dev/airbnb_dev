package controllers

import (
	"AuthInGo/services"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController { //constructor
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get user by id called in UserController ")
	uc.UserService.GetUserById()
	w.Write([]byte("user fetching endpoint"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create user by id called in UserController ")
	uc.UserService.CreateUser()
	w.Write([]byte("user fetching endpoint"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login user by id called in UserController ")
	uc.UserService.LoginUser()
	w.Write([]byte("user fetching endpoint"))
}