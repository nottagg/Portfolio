package main

import (
	"context"
	"log/slog"
	"os"
)

// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

func main() {
	//Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Info("Connecting to database")
	db := DBInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
	conn, err := db.DBConnect()
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	logger.Info("Connected to database")

	server := ServerInfo{
		Port:   os.Getenv("SERVER_PORT"),
		Db:     conn,
		logger: logger,
	}
	logger.Info("Starting server")
	server.RunServer()
	logger.Info("Server started on ", "Port", server.Port)
}
