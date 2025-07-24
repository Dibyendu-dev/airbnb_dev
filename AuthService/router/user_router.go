package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middleware"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController)Router{
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Regiter(r chi.Router){
	
	r.With(middleware.JWTAuthMiddleware).Get("/profile",ur.userController.GetUserById)
	r.With(middleware.UserCreateRequestValidator).Post("/signup",ur.userController.CreateUser)
	r.With(middleware.UserLoginRequestValidator).Post("/login",ur.userController.LoginUser)

}