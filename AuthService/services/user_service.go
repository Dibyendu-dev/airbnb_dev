package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
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