// infrastructures/gofiber/fiber.go
package gofiber

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"eql/configs"
	"eql/internal/handlers/middleware"
	"eql/internal/infrastructures/database"
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
	conf, err := configs.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	// สร้าง connection pool สำหรับ db1 และ db2
	dbConnections, err := database.NewDBConnections(*conf)
	if err != nil {
		panic(err)
	}
	// ส่งค่า connection pool ของฐานข้อมูลไปยัง Context ในที่นี้คือ db1 และ db2
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("main", dbConnections.DB1)
		c.Locals("db2", dbConnections.DB2)
		return c.Next()
	})
	// c := app.AcquireCtx(&fasthttp.RequestCtx{})
	// defer app.ReleaseCtx(c)

	// dbValue := c.Locals("main")
	// if dbValue == nil {
	// 	// หากไม่พบค่า db1 ที่ต้องการให้ทำการจัดการข้อผิดพลาดตามที่คุณต้องการ
	// 	log.Fatalf("Error: failed to get db1 from context")
	// }

	// db1, ok := dbValue.(*gorm.DB)
	// if !ok {
	// 	// หากไม่สามารถแปลงค่าให้เป็น *gorm.DB ได้ให้ทำการจัดการข้อผิดพลาดตามที่คุณต้องการ
	// 	log.Fatalf("Error: failed to convert db1 to *gorm.DB")
	// }
	// _ = db1
	// if db1 == nil {
	// 	// หากไม่พบค่า db1 ที่ต้องการให้ทำการจัดการข้อผิดพลาดตามที่คุณต้องการ
	// 	log.Fatalf("Error: failed to get db1 from context")
	// }

	// user := []models.User{}
	// err = db1.Find(&user).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user)

	app.Use(
		logger.New(middleware.LoggerFormat()), // เพิ่ม Logger Middleware เพื่อแสดงข้อมูลการทำงานของ API
		middleware.TimeLogger,                 // เพิ่ม middleware TimeLogger เพื่อบันทึกเวลาที่ใช้ในการ request
		middleware.RequestLoggerMiddleware,    //Add RequestLoggerMiddleware to log requests and responses in JSON format
	)
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

//		db1 := c.Locals("db1").(*gorm.DB)
//		db2 := c.Locals("db2").(*gorm.DB)
//การแปลง *fiber.App เป็น ctx
// ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
// err := CheckDatabaseConnection()
// app.ReleaseCtx(ctx)
