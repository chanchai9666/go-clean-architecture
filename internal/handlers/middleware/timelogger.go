package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func TimeLogger(c *fiber.Ctx) error {
	start := time.Now()

	if err := c.Next(); err != nil {
		return err
	}

	elapsed := time.Since(start)
	c.Append("X-Response-Time", elapsed.String())

	// Append the elapsed time to the context locals for logging in the console
	c.Locals("Elapsed", elapsed.String())

	return nil
}
