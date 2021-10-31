package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Colour    string
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured: %s", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_ENDPOINT"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Failed to migrate DB")
	}

	db.Create(&User{FirstName: "John", LastName: "Doe", Email: "john@mail.com", Password: "password", Colour: "#efefef"})

	var user User
	db.First(&user, "FirstName = ?", "John")

	db.Model(&user).Update("Email", "hello@mail.com")

	// Setup Form-User relationship (1-many r/ship)
	// Setup web server to use CRUD
	// Test queries
}
