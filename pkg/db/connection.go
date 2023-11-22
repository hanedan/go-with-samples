package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Connect(ctx context.Context) (*sql.DB, error) {
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, fmt.Errorf("DB_HOST environment variable required but not set")
	}
	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, fmt.Errorf("DB_PORT environment variable required but not set")
	}
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, fmt.Errorf("DB_USER environment variable required but not set")
	}
	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		return nil, fmt.Errorf("DB_PASS environment variable required but not set")
	}
	database, ok := os.LookupEnv("DB_DATABASE")
	if !ok {
		return nil, fmt.Errorf("DB_DATABASE environment variable required but not set")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, database)
	// driver is not "sqlserver". it's "postgres"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, db.PingContext(ctx)
}
