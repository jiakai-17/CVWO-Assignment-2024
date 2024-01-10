package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

var conn *pgx.Conn = nil
var err error = nil

func GetConnection() *pgx.Conn {
	ctx := context.Background()

	connString := os.Getenv("DATABASE_URL")

	if connString == "" {
		log.Println("No DATABASE_URL environment variable found, using dev database")
		connString = "host=localhost user=postgres dbname=cvwo-1 password=cs2102 port=5432"
	}

	conn, err = pgx.Connect(ctx, connString)

	if err != nil {
		log.Fatal("[ERROR] Unable to connect to database: ", err)
		return nil
	}
	return conn
}
