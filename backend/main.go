package main

import (
	"fmt"
	"log"
	"net/http"
)

// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

func main() {
    http.HandleFunc("/all", handleAll)
	fmt.Print("Listening on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))

}