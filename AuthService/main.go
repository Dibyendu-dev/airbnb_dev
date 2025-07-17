package main

import (
	"AuthInGo/app"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	cfg := app.Config{
		Addr: ":3000",
	}

	app := app.Application{
		Config:  cfg,
	}

	app.Run()
}


