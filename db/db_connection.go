package db_connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

type DB = pgxpool.Pool

// Use pgxpool for concurrent connections
// if you have multiple threads working with a DB at the same time, you must use pgxpool
func Init(
	username string,
	password string,
	host string,
	port string,
	database string,
) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	var err error

	db, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connected to database")
}

func Close() {
	db.Close()
}

func GetDB() *pgxpool.Pool {
	return db
}
