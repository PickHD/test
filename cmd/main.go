package main

import (
	"fmt"
	"test/internal/application"
	"test/internal/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app, err := application.NewApplication()
	if err != nil {
		panic(err)
	}

	http := infrastructure.InitializeRouter(app)

	http.Listen(fmt.Sprintf(":%d", app.Config.General.AppPort))
}
