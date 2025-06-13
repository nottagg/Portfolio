package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	// Extract username from JWT claims
	authHeader := r.Header.Get("Authorization")
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get user_id from username
	var userID int
	err := h.db.QueryRow(context.Background(), "SELECT id FROM users WHERE username=$1", username).Scan(&userID)
	if err != nil {
		h.logger.Error("Failed to get user ID", "error", err)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	rows, err := h.db.Query(context.Background(), `
		SELECT id, name, description, due_date, status
		FROM tasks
		WHERE user_id=$1
	`, userID)
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
		if err := rows.Scan(&item.ID, &item.Name, &item.Description, &dueDate, &status); err != nil {
			h.logger.Error("Failed to scan task row", "error", err)
			continue
		}
		item.UserID = userID
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

var jwtSecret = []byte("your_secret_key") // Use a secure key in production

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// For demo: accept any username/password from JSON body
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// TODO: Validate credentials from DB
	if creds.Username == "" || creds.Password == "" {
		http.Error(w, "Missing credentials", http.StatusUnauthorized)
		return
	}
	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": creds.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func jwtAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
