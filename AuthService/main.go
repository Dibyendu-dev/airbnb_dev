package main

import (
	"AuthInGo/app"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	cfg := app.NewConfig(":3000")

	app := app.NewApplication(cfg)

	app.Run()
}


