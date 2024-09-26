// package main

// import (
// 	"fmt"

// 	"github.com/gofiber/fiber/v2"
// )

// func main() {
// 	app := fiber.New()

// 	app.Get("/hello", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello World")
// 	})

// 	err := app.Listen(":8082")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
