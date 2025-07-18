package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	router :=chi.NewRouter()
	
	//routes
	router.Get("/ping",controllers.PingHandler)


	return  router
}