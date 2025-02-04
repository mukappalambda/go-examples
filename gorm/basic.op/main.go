package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID int    `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("error migrating: %v\n", err)
	}
	fmt.Println("creating a single user")
	db.Create(&User{UserID: 1, Name: "alpha"})
	users := []User{
		{
			UserID: 2,
			Name:   "beta",
		},
		{
			UserID: 3,
			Name:   "gamma",
		},
	}
	fmt.Println("creating multiple users")
	db.Create(&users)

	fmt.Println("reading a single user")
	var user User
	db.First(&user, 1)
	fmt.Println(user)
	fmt.Println("reading multiple users")
	foundUsers := make([]User, 0)
	db.Find(&foundUsers)
	for _, user := range foundUsers {
		fmt.Printf("found user: %+v\n", user)
	}

	// fmt.Println("Update ...")
	// db.Model(&user).Where("Name = ?", "alex").Update("Name", "bob")
	// db.First(&user, 1)
	// fmt.Println(user)

	fmt.Println("Delete ...")
	db.Where("Name = ?", "bob").Delete(&user)
	fmt.Println(user)
}
