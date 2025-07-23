package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
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
	var payload dto.LoginUserRequestDTO
	if jsonErr := utils.ReadJsonBody(r,&payload); jsonErr!= nil{
		w.Write([]byte("something went wrong while loging in"))
		return
	}
	
	// check validation
	if validationErr := utils.Validator.Struct(payload); validationErr!= nil {
		w.Write([]byte("invalid input data"))
		return
	}


	jwtToken,err :=uc.UserService.LoginUser()
	if err!=nil{
		w.Write([]byte("something went wrong while loging in"))
		return
	}
	 response := map[string]any{
		"message": "user logged in successfully",
		"data": jwtToken,
		"success": true,
		"error":nil,
	 }
	 utils.WriteJsonResponse(w,http.StatusOK,response)

	
}