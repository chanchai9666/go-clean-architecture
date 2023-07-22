// infrastructures/gofiber/fiber.go
package gofiber

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
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

	// ตัวแปร option เพื่อกำหนดการแสดงข้อมูลใน Logger
	option := logger.Config{
		Format:     "${time} ${method} ${path} - ${ip} - ${status} - ${locals:Elapsed}\n", // Use ${locals:Elapsed} to print the elapsed time
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     os.Stdout,
	}
	// เพิ่ม Logger Middleware เพื่อแสดงข้อมูลการทำงานของ API โดยใช้ตัวเลือกที่กำหนดไว้ใน option
	app.Use(logger.New(option))
	// เพิ่ม middleware TimeLogger เพื่อบันทึกเวลาที่ใช้ในการ request
	app.Use(middleware.TimeLogger)

	return app
}
