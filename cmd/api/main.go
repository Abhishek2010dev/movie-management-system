package main

import (
	"log"

	"github.com/Abhishek2010dev/movie-management-system/config"
	"github.com/Abhishek2010dev/movie-management-system/server"
)

func main() {
	cfg := config.New()
	server := server.New(cfg)
	log.Printf("Server is running at %s:%s", cfg.Server.Host, cfg.Server.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
