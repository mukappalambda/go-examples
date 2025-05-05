package main

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	dsn := "user=postgres password=password dbname=demo host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	if err := db.Callback().Query().Register("gorm:mycallback", func(*gorm.DB) {
		fmt.Println("query callback:", time.Now())
	}); err != nil {
		return fmt.Errorf("failed to register callback: %s", err)
	}
	var result string
	db.First(&result)
	return nil
}
