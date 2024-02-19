package postgres

import (
	"context"
	"database/sql"
	"fmt"

	// Import your preferred PostgreSQL driver (e.g., lib/pq)
	_ "github.com/lib/pq" // Example assuming pq driver
)

const (
	host     = "localhost"
	port     = 5432 // Standard PostgreSQL port
	user     = "postgres"
	password = "1234"
	database = "postgres"
)

// Connect creates a connection to the PostgreSQL database.
func Connect(ctx context.Context) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, database)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return db, nil
}

// Close closes the database connection.
func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		fmt.Printf("error closing database connection: %v\n", err)
	}
}
