package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middleware"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}


func SetupRouter(UserRouter Router, RoleRouter Router) *chi.Mux {
	chirouter :=chi.NewRouter()
	chirouter.Use(middleware.RequestLogger)
	chirouter.Use(middleware.RateLimitMiddleware)
	//routes
	chirouter.Get("/ping",controllers.PingHandler)
	chirouter.HandleFunc("/fakestoreservice/*",utils.ProxyToService("https://fakestore.api.in","/fakestoreservice"))
	UserRouter.Register(chirouter)
	RoleRouter.Register(chirouter)
	return  chirouter
}

// http://localhost:3001/fakestoreservice/products