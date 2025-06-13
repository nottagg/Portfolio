package main

import (
	"log/slog"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type Handler struct {
	db     *pgx.Conn
	logger *slog.Logger
}

func (h *Handler) handleEmpty(w http.ResponseWriter, r *http.Request) {
	// Handle empty request
	h.logger.Info("Received an empty request")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
	h.logger.Info("Handled empty request")
}
func (h *Handler) handleGetTask(w http.ResponseWriter, r *http.Request) {
	// Handle GET task request
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("GET task"))
}
