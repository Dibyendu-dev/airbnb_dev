package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/router"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string  //port
}
type Application struct {
	Config Config
}

func NewConfig() Config {  //constructor fn
	
	port:= config.GetString("PORT",":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) *Application {
	return  &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server on",app.Config.Addr)
	return server.ListenAndServe()
}
