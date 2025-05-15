package main

import (
	"github.com/Abhishek2010dev/movie-management-system/server"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	server := server.New()
	app := server.Setup()
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
