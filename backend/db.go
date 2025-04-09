package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DBInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (db *DBInfo) DBConnect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), db.GetConnectionString())
}

func (db *DBInfo) GetConnectionString() string {
	return "host=" + db.Host +
		" port=" + db.Port +
		" user=" + db.User +
		" password=" + db.Password +
		" dbname=" + db.Database
}
