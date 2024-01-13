package database

import (
	"backend/utils"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"os"
)

var conn *pgx.Conn = nil
var err error = nil

// GetConnection Returns a connection to the database. Returns nil if unable to connect.
// Remember to close the connection by calling conn.Close(ctx).
func GetConnection() *pgx.Conn {
	ctx := context.Background()

	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		utils.Log("database", "No database URL provided!", errors.New("no database URL provided"))
	}

	conn, err = pgx.Connect(ctx, connString)

	if err != nil {
		utils.Log("database", "Unable to connect to database", err)
		return nil
	}
	return conn
}

// CloseConnection Closes a connection to the database.
func CloseConnection(conn *pgx.Conn) {
	ctx := context.Background()
	err := conn.Close(ctx)
	if err != nil {
		utils.Log("database", "Unable to close connection", err)
		return
	}
}
