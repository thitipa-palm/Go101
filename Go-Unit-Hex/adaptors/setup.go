package adaptors

import (
	"github.com/gofiber/fiber/v2"
)

func SetUp(orderHandler *HttpOrderHandler) *fiber.App {
	app := fiber.New()
	app.Post("/order", orderHandler.CreateOrder)
	return app
}
