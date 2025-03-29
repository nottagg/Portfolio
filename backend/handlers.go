package main

import (
	"net/http"
)

func handleAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//TODO: Return all todo list items
	case http.MethodDelete:
		//TODO: Delete all todo list items
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
