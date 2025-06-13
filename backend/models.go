package main

import (
	"time"
)

type ItemStatus string

const (
	StatusPending    ItemStatus = "Pending"
	StatusInProgress ItemStatus = "In Progress"
	StatusCompleted  ItemStatus = "Completed"
	StatusCancelled  ItemStatus = "Cancelled"
)

type User struct {
	ID       int
	Username string
	Password string // Store hashed passwords in production
}

type ToDoItem struct {
	ID          int
	UserID      int
	Name        string
	Description string
	DueDate     time.Time
	Status      ItemStatus
}
