package main

import (
	"fmt"
	"log"

	"database/sql"

	_ "github.com/lib/pq"
)

type Todo struct {
	Name        string
	IsCompleted bool
}

func main() {
	dataSourceName := "postgres://postgres:password@localhost/my_db?sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal()
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO todo (name, is_completed) VALUES ($1, $2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	todos := []Todo{
		{Name: "Buy groceries", IsCompleted: false},
		{Name: "Do laundry", IsCompleted: false},
		{Name: "Finish project", IsCompleted: true},
	}

	for _, todo := range todos {
		fmt.Println(todo)
		_, err = stmt.Exec(todo.Name, todo.IsCompleted)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Todo records inserted successfully.")
}
