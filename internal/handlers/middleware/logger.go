package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rs/zerolog/log"
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

func RequestLoggerMiddleware(c *fiber.Ctx) error {
	// Call next handler
	err := c.Next()

	// Get the status code
	status := c.Response().StatusCode()

	// Get the request URL
	url := c.Request().URI().String()

	// Get the request method
	method := c.Request().Header.Method()

	// Get the response body as a string
	body := c.Response().Body()

	// Log the request and response using Zerolog
	log.Info().
		Int("status", status).
		Str("url", url).
		Str("method", string(method)).
		Str("body", string(body)).
		Msg("Request and Response")

	return err
}

func LoggerFormat() logger.Config {
	//กำหนดการแสดงข้อมูลใน Logger
	return logger.Config{
		Format:     "${time} ${method} ${path} - ${ip} - ${status} - ${locals:Elapsed}\n", // Use ${locals:Elapsed} to print the elapsed time
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     os.Stdout,
	}
}
