package main

import (
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

type ItemStatus string

const (
	StatusPending    ItemStatus = "Pending"
	StatusInProgress ItemStatus = "In Progress"
	StatusCompleted  ItemStatus = "Completed"
	StatusCancelled  ItemStatus = "Cancelled"
)

type ToDoItem struct {
	Name        string
	Description string
	DueDate     time.Time
	Status      ItemStatus
}

type DBInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type ServerInfo struct {
	Addr   string
	Db     *pgx.Conn
	logger *slog.Logger
}
