package main

import (
	"fmt"
	"log"
	"net/http"
)

// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

func Get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
    http.HandleFunc("/", Get)
	fmt.Print("Listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}