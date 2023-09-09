package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId int    `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Create ...")
	db.AutoMigrate(&User{})
	db.Create(&User{UserId: 1, Name: "alex"})

	fmt.Println("Read ...")
	var user User
	db.First(&user, 1)
	fmt.Println(user)

	// fmt.Println("Update ...")
	// db.Model(&user).Where("Name = ?", "alex").Update("Name", "bob")
	// db.First(&user, 1)
	// fmt.Println(user)

	fmt.Println("Delete ...")
	db.Where("Name = ?", "bob").Delete(&user)
	fmt.Println(user)

}
