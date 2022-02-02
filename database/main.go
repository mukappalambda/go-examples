package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Id int64
	Username string
	Score float32
}

func main() {
	// get a database handle and connect
	connStr := "postgres://postgres:postgres@postgres/demo?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected!")

	// I have already insert data into mytable
	
	fmt.Println("query...")
	
	// select * from mytable
	users, _ := GetUsers(db)
	fmt.Println("users:", users)
	
	// select * from mytable (multiple)
	users1, _ := GetUser(db, "alex")
	fmt.Println("user1:", users1)
	
	// select * from mytable (single)
	singleUser, _ := GetSingleUser(db, "bob")
	fmt.Println("single user:", singleUser)

	// add a new user
	newUser := User{
		Id: 3,
		Username: "joe",
		Score: 30.3,
	}
	num, _ := AddUser(db, newUser)
	fmt.Println("number:", num)
	otherUsers, _ := GetUsers(db)
	fmt.Println("query users again:", otherUsers)
}

func GetUsers(db *sql.DB) ([]User, error) {
	var users []User
	var user User

	rows, err := db.Query("SELECT * FROM mytable")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Score); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(db *sql.DB, name string) ([]User, error) {
	var users []User
	var user User

	rows, err := db.Query(`SELECT * FROM mytable WHERE username = $1`, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Username, &user.Score); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetSingleUser(db *sql.DB, name string) (User, error) {
	var user User

	rows := db.QueryRow(`SELECT * FROM mytable WHERE username = $1`, name)

	if err := rows.Scan(&user.Id, &user.Username, &user.Score); err != nil {
		return user, err
	}

	return user, nil
}

func AddUser(db *sql.DB, user User) (int64, error) {
	result, err := db.Exec(`INSERT INTO mytable (id, username, score) values ($1, $2, $3)`, user.Id, user.Username, user.Score)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}