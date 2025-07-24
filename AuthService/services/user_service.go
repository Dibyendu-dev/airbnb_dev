package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById(id string)  (*models.User ,error)
	CreateUser(payload *dto.CreateUserRequestDTO) (*models.User ,error)
	LoginUser(payload *dto.LoginUserRequestDTO)	(string,error)
	
}

type UserServiceImpl struct {
	userRepository db.UserRepository //this depends on interface ,not userServiceImpl(class) *DI
}

func NewUserService(_userRepository db.UserRepository) UserService { //constructor fn
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById(id string) (*models.User ,error) {
	fmt.Println("fetching user in user service")
	users,err :=u.userRepository.GetUserById(id)
	if err != nil{
		fmt.Println("error fetching user",err)
		return nil,err
	}
	return users, nil
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User ,error){
	fmt.Println("creating user in userService")

	// step 1: hash the password
	hashedPassword,err := utils.HashPassword(payload.Password)
	
	if err != nil{
		fmt.Println("error to hashed password",err)
		return nil,err
	}

	// step 2: call repository layer to create a user
	user,err := u.userRepository.Create(payload.Username,payload.Email,hashedPassword)
	if err != nil{
		fmt.Println("error to create password",err)
		return nil,err
	}
	return user,nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string,error){
	
	email:= payload.Email
	password := payload.Password

	//step 1: make a repository call to get the user by email
	user,err :=u.userRepository.GetUserByEmail(email)
	if err != nil{
		fmt.Println("error to find users",err)
		return "",err
	}
	// step 2:user exsist or not
	if user == nil{
		fmt.Println("user not found")
		return "",fmt.Errorf("no user found with email")
	}

	// step 3: if user exsists check password with hashpassword
	isPasswordValid := utils.CheckPasswordHash(password, user.Password)

	if !isPasswordValid {
		fmt.Println("password doesn't matched ")
		return "",nil
	}

	// step 4: if password match, sent jwt token
	jwtPayload:= jwt.MapClaims{
		"email": user.Email,
		"id":user.Id,
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwtPayload) 
	tokenString,err :=token.SignedString([]byte(env.GetString("JWT_SECRET","token")))
	 if err!= nil{
		fmt.Println("error signing token",err)
		return "", err
	 }
	 fmt.Println("JWT token:",tokenString)
	return tokenString ,nil
}