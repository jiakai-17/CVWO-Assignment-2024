package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

var conn *pgx.Conn = nil
var err error = nil

func Init() {
	ctx := context.Background()
	conn, err = pgx.Connect(ctx, "user=postgres dbname=cvwo-1 password=cs2102")
	if err != nil {
		log.Fatal("[ERROR] Unable to connect to database: ", err)
		return
	}
	defer conn.Close(ctx)
}

func GetConnection() *pgx.Conn {
	return conn
}
