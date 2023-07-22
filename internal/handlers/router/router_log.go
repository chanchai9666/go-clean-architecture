package router

import (
	"github.com/gofiber/fiber/v2"
)

func LogRouter(router fiber.Router) {
	// Define the routes for /logs using the provided router
	router.Get("/list", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET Log request!")
	})
}
