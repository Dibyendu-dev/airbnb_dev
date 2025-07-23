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
	
	var payload dto.LoginUserRequestDTO
	if jsonErr := utils.ReadJsonBody(r, &payload); jsonErr != nil {
				utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"something went wrong while loggin in",jsonErr)

		return
	}

	// check validation
	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"invalid input data ",validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w,http.StatusInternalServerError,"failed to login user ",err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "user logged in successfully", jwtToken)

}
