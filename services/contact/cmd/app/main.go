package main

import (
	"architecture_go/pkg/store/postgres"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"architecture_go/services/contact/internal/delivery"
	"architecture_go/services/contact/internal/repository"
	_ "architecture_go/services/contact/internal/useCase"
)

func main() {
	ctx := context.Background()

	db, err := postgres.Connect(ctx)
	if err != nil {
		fmt.Println("failed to connect to database:", err)
		return
	}

	defer postgres.Close(db)

	// Use the database connection here, e.g.:
	rows, err := db.Query("SELECT * FROM your_table")
	if err != nil {
		fmt.Println("error querying database:", err)
		return
	}

	defer rows.Close()

	// Process rows...

	fmt.Println("Connected to PostgreSQL successfully!")

	// Parse flags
	var (
		port = flag.Int("port", 8080, "The port to listen on")
	)
	flag.Parse()

	// Create database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
	if err != nil {
		log.Fatal(err)
	}

	// Create repositories
	repo := repository.NewContactRepository(db)
	groupRepo := repository.NewGroupRepository(db)

	// Create use case
	usecase := usecase.NewContactUseCase(repo, groupRepo)

	// Create HTTP adapter
	adapter := delivery.NewHTTPAdapter(usecase)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + string(*port),
		Handler: adapter.Router(),
	}

	// Start HTTP server
	go func() {
		log.Println("Starting HTTP server on port", *port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Shutdown HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("HTTP server stopped")
}
