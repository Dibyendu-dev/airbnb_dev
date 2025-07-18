package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Regiter(r chi.Router)
}


func SetupRouter(UserRouter Router) *chi.Mux {
	chirouter :=chi.NewRouter()
	
	//routes
	chirouter.Get("/ping",controllers.PingHandler)
	UserRouter.Regiter(chirouter)

	return  chirouter
}