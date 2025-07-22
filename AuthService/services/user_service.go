package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
	env "AuthInGo/config/env"
	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser()	(string,error)
	
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

func (u *UserServiceImpl) LoginUser() (string,error){
	
	email:= "user2@test.com"
	password := "hashuser2pass123"

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
	payload:= jwt.MapClaims{
		"email": user.Email,
		"id":user.Id,
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,payload) 
	tokenString,err :=token.SignedString([]byte(env.GetString("JWT_SECRET","token")))
	 if err!= nil{
		fmt.Println("error signing token",err)
		return "", err
	 }
	 fmt.Println("JWT token:",tokenString)
	return tokenString ,nil
}