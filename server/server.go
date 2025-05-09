package server

import (
	"fmt"
	"net/http"

	"github.com/Abhishek2010dev/movie-management-system/config"
)

type Server struct {
	cfg config.Config
}

func New(cfg config.Config) *http.Server {
	server := Server{cfg}

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      server.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
}
