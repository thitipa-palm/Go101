package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/thitipa-palm/go-Unit-Hex/adaptors"
	"github.com/thitipa-palm/go-Unit-Hex/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {

	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			// IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			// ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful: true, // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database") // break or kill process
	}
	// print(db)
	fmt.Println("Connection successful")

	// secondary => primary
	orderRepo := adaptors.NewGormRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adaptors.NewHttpOrderHandler(orderService)

	app := adaptors.SetUp(orderHandler)
	db.AutoMigrate(&core.Order{})

	app.Listen(":8081")
}
