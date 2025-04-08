package main

import (
	"net/http"
)

func (s *ServerInfo) handleEmpty(w http.ResponseWriter, r *http.Request) {
	// Handle empty request
	s.logger.Info("Received an empty request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
	s.logger.Info("Handled empty request")
}
func handleGetTask(s ServerInfo, w http.ResponseWriter, r *http.Request) {
	// Handle GET task request
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET task"))
}
