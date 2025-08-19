package router

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetUpRouter(ReviewRouter Router) *chi.Mux {

	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)

	ReviewRouter.Register(chiRouter)

	return chiRouter
}
