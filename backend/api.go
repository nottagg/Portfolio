package main

import (
	"net/http"
)

type API struct {
	addr string
}

func APIConstructor(addr string) *API {
	return &API{addr: addr}
}

func RunServer(api *API) error {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handleGet)
	server := http.Server{
		Addr:    api.addr,
		Handler: router,
	}
	return server.ListenAndServe()
}
