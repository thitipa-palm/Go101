package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

// InitializeDB initializes the database and automigrates the User model.
func InitializeDB() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "myuser"
		password = "mypassword"
		dbname   = "mydatabase"
	)
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})
	return db
}

// AddUser adds a new user to the database.
func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := User{Fullname: fullname, Email: email, Age: age}

	// Check if email already exists
	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	// Save the new user
	result := db.Create(&user)
	return result.Error
}

func main() {
	db := InitializeDB()
	// Your application code
	err := AddUser(db, "Tiger Palm2", "tiger.palm2@example.com", 24)
	if err != nil {
		log.Println("Error!!!", err)
	}
}
