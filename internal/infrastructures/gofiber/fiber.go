// infrastructures/gofiber/fiber.go
package gofiber

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"eql/configs"
	"eql/internal/handlers/middleware"
	"eql/internal/infrastructures/database"
)

// NewServer สร้างเซิร์ฟเวอร์ GoFiber ใหม่
func NewServer() *fiber.App {

	// โหลด config จากไฟล์ config.yaml
	conf, err := configs.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

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

	// สร้าง connection pool สำหรับ db1
	dsnMain := database.GenerateDSN(conf.DatabaseConfig, "Main")
	db1, err := database.NewDBConnection(dsnMain)
	if err != nil {
		// ตรวจสอบและจัดการ error ที่เกิดขึ้นในกรณีเชื่อมต่อฐานข้อมูลไม่สำเร็จ
		panic(fmt.Errorf("failed to connect to db1: %w", err))
	}

	// สร้าง connection pool สำหรับ db2
	dsnMain2 := database.GenerateDSN(conf.DatabaseConfig, "Main2")
	db2, err := database.NewDBConnection(dsnMain2)
	if err != nil {
		// ตรวจสอบและจัดการ error ที่เกิดขึ้นในกรณีเชื่อมต่อฐานข้อมูลไม่สำเร็จ
		panic(fmt.Errorf("failed to connect to db2: %w", err))
	}

	// ส่งค่า connection pool ของฐานข้อมูลไปยัง Context ในที่นี้คือ db1 และ db2
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db1", db1)
		c.Locals("db2", db2)
		return c.Next()
	})
	// การเชื่อมต่อสำเร็จ
	fmt.Println("Connected to db1 and db2 successfully!")
	return app
}
