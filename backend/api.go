package main

import (
	"net/http"
)

func RunServer(s *ServerInfo) error {
	router := http.NewServeMux()

	router.HandleFunc("GET /", s.handleEmpty)
	server := http.Server{
		Addr:    s.Addr,
		Handler: router,
	}
	return server.ListenAndServe()
}
