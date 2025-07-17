package app

import (
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

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      nil, //TODO: set up a chi router
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server on",app.Config.Addr)
	return server.ListenAndServe()
}
