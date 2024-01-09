package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

var conn *pgx.Conn = nil
var err error = nil

func GetConnection() *pgx.Conn {
	ctx := context.Background()
	conn, err = pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to database: ", err)
		return nil
	}
	return conn
}
