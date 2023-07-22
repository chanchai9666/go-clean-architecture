package router

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UsersRouter(router fiber.Router) {
	// Define the routes for /users using the provided router
	router.Get("/list", func(c *fiber.Ctx) error {
		fmt.Println("Start")
		time.Sleep(3 * time.Second)
		fmt.Println("End")
		return c.SendString("I'm a GET USER request!")
	})
}
