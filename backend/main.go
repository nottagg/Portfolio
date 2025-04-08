package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

func main() {
	//Logging
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//Database connection
	DBInfo := DBInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		DBInfo.Host,
		DBInfo.Port,
		DBInfo.User,
		DBInfo.Password,
		DBInfo.Database)
	// Connect to the database
	logger.Info("Connecting to database")
	conn, err := pgx.Connect(context.Background(), dbConnectionString)
	if err != nil {
		logger.Error("Unable to connect to database", "error", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	logger.Info("Connected to database")

	server := ServerInfo{
		Addr:   os.Getenv("SERVER_ADDR"),
		Db:     conn,
		logger: logger,
	}
	logger.Info("Starting server")
	fmt.Printf("Starting server on port " + server.Addr)
	RunServer(&server)
	logger.Info("Server started", "addr", server.Addr)

}
