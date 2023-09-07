// infrastructures/gofiber/fiber.go
package gofiber

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"eql/internal/handlers/middleware"
)

// NewServer สร้างเซิร์ฟเวอร์ GoFiber ใหม่
func NewServer() *fiber.App {

	app := fiber.New(fiber.Config{
		// Concurrency เป็นจำนวน worker threads ที่ใช้ในการจัดการ concurrent connections
		Concurrency:  10,                // ปรับค่าตามความเหมาะสมของเซิร์ฟเวอร์
		ReadTimeout:  10 * time.Second,  // ระยะเวลาที่เซิร์ฟเวอร์รอในการอ่าน request body
		WriteTimeout: 10 * time.Second,  // ระยะเวลาที่เซิร์ฟเวอร์รอในการเขียน response
		IdleTimeout:  120 * time.Second, // ระยะเวลาที่ connection ต้อง idle ก่อนถูกปิด
	})

	// โหลด config จากไฟล์ config.yaml
	// conf, err := configs.LoadConfig()
	// if err != nil {
	// 	panic(fmt.Errorf("failed to load config: %w", err))
	// }

	// กำหนดค่า CORS ดังนี้
	app.Use(cors.New())
	app.Use(
		logger.New(middleware.LoggerFormat()), // เพิ่ม Logger Middleware เพื่อแสดงข้อมูลการทำงานของ API
		middleware.TimeLogger,                 // เพิ่ม middleware TimeLogger เพื่อบันทึกเวลาที่ใช้ในการ request
		middleware.RequestLoggerMiddleware,    //Add RequestLoggerMiddleware to log requests and responses in JSON format
	)
	// สร้าง middleware สำหรับการบีบอัด
	app.Use(compress.New())
	//Recover กรณี panic หรือ มีเหตุให้ API หยุดทำงาน
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				sentry.CurrentHub().Recover(r)
			}
		}()
		return c.Next()
	})
	// OnRequest Hook
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("OnRequest Hook - Before processing the request")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// go database.CheckDatabaseStatus()

	return app
}
