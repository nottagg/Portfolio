package main

import (
	"fmt"
)

// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

func main() {
	server := APIConstructor(":8080")
	fmt.Printf("Starting server on port " + server.addr)
	RunServer(server)

}
