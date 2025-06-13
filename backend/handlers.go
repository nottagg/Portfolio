package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

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
	rows, err := h.db.Query(context.Background(), `
		SELECT name, description, due_date, status
		FROM tasks
	`)
	if err != nil {
		h.logger.Error("Failed to query tasks", "error", err)
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []ToDoItem
	for rows.Next() {
		var item ToDoItem
		var dueDate time.Time
		var status string
		if err := rows.Scan(&item.Name, &item.Description, &dueDate, &status); err != nil {
			h.logger.Error("Failed to scan task row", "error", err)
			continue
		}
		item.DueDate = dueDate
		item.Status = ItemStatus(status)
		tasks = append(tasks, item)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		h.logger.Error("Failed to encode tasks to JSON", "error", err)
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
	}
}
