package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var numFakeUsers = flag.Int("num-fake-users", 0, "number of fake users to create")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	dsn := "postgresql://postgres:password@localhost:5432/demo?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	ctx := context.Background()
	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %s", err)
	}

	if err := initDB(ctx, db); err != nil {
		return err
	}

	if err := createFakeUsers(ctx, db, *numFakeUsers); err != nil {
		return err
	}

	users, err := getAllUsers(db)
	if err != nil {
		return err
	}
	fmt.Println("all users:")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
	return nil
}

func initDB(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `drop table if exists users;`)
	if err != nil {
		return fmt.Errorf("failed to drop users table: %s", err)
	}
	_, err = db.ExecContext(ctx, `create table if not exists users (id serial primary key, name text not null, email text not null);`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %s", err)
	}
	return nil
}

func createFakeUsers(ctx context.Context, db *sql.DB, n int) error {
	_, err := db.ExecContext(ctx, `with t as (select left(gen_random_uuid()::text, 4) as n from generate_series(1, $1)) insert into users (name, email) select 'user_' || n::text as name, n::text || '@example.com' as email from t;`, n)
	if err != nil {
		return fmt.Errorf("failed to create fake users: %s", err)
	}
	return nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`select id, name, email from users;`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all users: %s", err)
	}
	defer rows.Close()
	users := make([]User, 0)
	var user User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to scan the current row: %s", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("got error when iterating rows: %s", err)
	}
	return users, nil
}
