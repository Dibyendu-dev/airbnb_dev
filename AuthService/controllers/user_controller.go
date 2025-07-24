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

//GetUser
func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching in UserController ")

	// extract userid from url params
	userId := r.URL.Query().Get("id")

	if userId =="" {
		userId = r.Context().Value("userID").(string) // fallback to context if not in url
	}

	fmt.Println("user id from context or query",userId)

	if userId == ""{
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"userId is required",fmt.Errorf("missing userid"))
		return
	}
	
	user,err :=uc.UserService.GetUserById(userId)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "failed to fetch user ", err)
		return
	}

	if user == nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "user not found ", fmt.Errorf("user with ID not found"))
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "user fetched in successfully", user)
	fmt.Println("user fetched succesfully", user)

	
}

// signup
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	
	payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)
	fmt.Println("Payload received:", payload)

	user, err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "failed to login user ", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "user created successfully", user)
	fmt.Println("user created succesfully", user)

}

// login
func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)
	fmt.Println("Payload received:", payload)

	jwtToken, err := uc.UserService.LoginUser(&payload)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "failed to login user ", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "user logged in successfully", jwtToken)
	fmt.Println("user logged in succesfully, JWT Token:", jwtToken)

}
