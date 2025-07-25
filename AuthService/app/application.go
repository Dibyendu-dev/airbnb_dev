package app

import (
	config "AuthInGo/config/env"
	dbConfig "AuthInGo/config/db"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string  //port
}
type Application struct {
	Config Config
	// Store db.Storage
}

func NewConfig() Config {  //constructor for Config
	
	port:= config.GetString("PORT",":8080")
	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) *Application {  //constructor for Application
	return  &Application{
		Config: cfg,
		// Store: *db.NewStorage(),
	}
}

func (app *Application) Run() error {
	db,err :=dbConfig.SetupDB()
	if err !=nil{
		fmt.Println("error connecting to database",err)
		return err
	}

	ur:= repo.NewUserRepository(db)
	us:= services.NewUserService(ur)
	uc:= controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)
	
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("starting server on",app.Config.Addr)
	return server.ListenAndServe()
}
