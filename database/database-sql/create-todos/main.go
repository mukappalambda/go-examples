package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

type Todo struct {
	Name        string
	IsCompleted bool
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_DSN"))
	if err != nil {
		return fmt.Errorf("failed to open database: %s", err)
	}
	defer db.Close()
	todos := []Todo{
		{Name: "Buy groceries", IsCompleted: false},
		{Name: "Do laundry", IsCompleted: false},
		{Name: "Finish project", IsCompleted: true},
	}
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := addTodos(ctx, db, todos...); err != nil {
		return fmt.Errorf("failed to add todos: %s", err)
	}
	fmt.Println("Todo records inserted successfully.")
	return nil
}

func addTodos(ctx context.Context, db *sql.DB, todos ...Todo) error {
	stmt, err := db.PrepareContext(ctx, "INSERT INTO todo (name, is_completed) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %s", err)
	}
	defer stmt.Close()
	for _, todo := range todos {
		fmt.Println(todo)
		if _, err := stmt.ExecContext(ctx, todo.Name, todo.IsCompleted); err != nil {
			return fmt.Errorf("failed to execute statement: %s", err)
		}
	}
	return nil
}
