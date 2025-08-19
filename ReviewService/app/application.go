package app

import (
	dbConfig "ReviewService/config/db"
	config "ReviewService/config/env"
	"ReviewService/controllers"
	repo "ReviewService/db/repository"
	"ReviewService/services"
	"ReviewService/router"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func NewConfig() Config {
	port := config.GetString("PORT", ":8081")

	return Config{
		Addr: port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() error {
	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("error setting up databse:", err)
		return err
	}

	rr:= repo.NewReviewRepository(db)
	rs:= services.NewReviewService(rr)
	rc:= controllers.NewReviewController(rs)
	rRouter:= router.NewReviewRouter(rc)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(rRouter),
		ReadTimeout:  10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
	}

	fmt.Println("Starting review service on ", app.Config.Addr)

	return server.ListenAndServe()
}
