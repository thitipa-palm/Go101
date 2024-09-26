package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
}

func CreateBook(db *gorm.DB, book *Book) {
	result := db.Create(book)
	if result.Error != nil {
		log.Fatalf("Error creating book: %v", result.Error)
	}
	fmt.Println("Book created successfully")
}

func GetBook(db *gorm.DB, id uint) *Book {
	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		log.Fatalf("Error finding book: %v", result.Error)
	}
	return &book
}

func UpdateBook(db *gorm.DB, book *Book) {
	result := db.Save(book)
	if result.Error != nil {
		log.Fatalf("Error updating book: %v", result.Error)
	}
	fmt.Println("Book updated successfully")
}

func DeleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)
	if result.Error != nil {
		log.Fatalf("Error deleting book: %v", result.Error)
	}
	fmt.Println("Book deleted successfully")
}

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
	// Config Gorm เหมือนเดิม
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&Book{})
	fmt.Println("Database migration completed!")
	// Create a new book
	newBook := Book{Name: "The Go Programming Language", Author: "Palm", Description: "Comprehensive guide to Go"}
	CreateBook(db, &newBook)

	// // Get a book
	// book := GetBook(db, 1) // Assuming the ID of the book is 1
	// fmt.Println("Book Retrieved:", book)

	book := Book{Name: "The Go Programming Language, Updated Edition", Author: "Palm", Description: "Comprehensive guide to Go"}
	// // Update a book
	// book.ID = 1
	// book.Name = "The Go Programming Language, Updated Edition"
	UpdateBook(db, &book)

	// Delete a book
	// DeleteBook(db, book.ID)
}
