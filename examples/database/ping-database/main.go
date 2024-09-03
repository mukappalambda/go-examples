package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	dataSourceName := "postgres://postgres:password@localhost/my_db?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected.")
	fmt.Printf("max open connections: %d\n", db.Stats().MaxOpenConnections)
	// Database connected.
	// max open connections: 50
}
