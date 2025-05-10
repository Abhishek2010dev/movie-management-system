package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Abhishek2010dev/movie-management-system/config"
)

type Server struct {
	cfg config.Config
	db  *sql.DB
}

func New(cfg config.Config, db *sql.DB) *http.Server {
	server := Server{cfg, db}

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler:      server.RegisterRoutes(),
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}
}
