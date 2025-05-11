package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// BuildConnString returns the connection string in pgx/sql format
func BuildConnString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)
}

// Connect returns a *pgx.Conn for use in services
func ConnectDB() (*pgx.Conn, error) {
	connStr := BuildConnString()
return pgx.Connect(context.Background(), connStr)
}
