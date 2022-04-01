package main

import (
	"kurneo/cmd/server"
	"kurneo/config"
	"kurneo/internal/infrastructure/dbconn"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Go REST API
// @version 1.0
// @description Golang REST API
// @contact.name Giang Nguyen
// @contact.url https://github.com/kurneo
// @contact.email giangnguyen.neko.130@gmail.com

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

func main() {
	if err := run(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initialize() {
	// Read env file
	env, errReadEnv := godotenv.Read(".env")

	if errReadEnv != nil {
		panic(errReadEnv)
	}

	for key, value := range env {
		if errSet := os.Setenv(key, value); errSet != nil {
			panic(errSet)
		}
	}

	// Init DB
	config.NewDBConfig()
	_, err := dbconn.NewConnection()

	if err != nil {
		panic(err)
	}
}

func run() error {
	initialize()
	app := echo.New()

	return server.Start(app)
}
