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

type ToDoItem struct {
	Name        string
	Description string
	DueDate     time.Time
	Status      ItemStatus
}
