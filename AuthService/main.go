package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
	dbconfig "AuthInGo/Config/db"
)

func main() {
	config.Load()
	cfg := app.NewConfig()

	app := app.NewApplication(cfg)
	dbconfig.SetupDB()
	app.Run()
}


