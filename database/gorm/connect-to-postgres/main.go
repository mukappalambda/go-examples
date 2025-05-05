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
		return fmt.Errorf("failed to open db: %s", err)
	}
	var result time.Time
	db.Raw("select now();").Scan(&result)
	var configFile string
	db.Raw("show config_file;").Scan(&configFile)
	fmt.Println("now:", result)
	fmt.Println("config file:", configFile)
	return nil
}
