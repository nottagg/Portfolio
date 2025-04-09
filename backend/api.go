package main

import (
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type ServerInfo struct {
	Port   string
	Db     *pgx.Conn
	logger *slog.Logger
}

func (s *ServerInfo) RunServer() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /", s.handleEmpty)
	server := http.Server{
		Addr:    ":" + s.Port,
		Handler: router,
	}
	return server.ListenAndServe()
}
