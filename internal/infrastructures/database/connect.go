package database

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"eql/configs"
)

func GetMainDatabase(c *fiber.Ctx) *gorm.DB {
	dbValue := c.Locals("main") //db หลัก
	if dbValue == nil {
		dsnMain := GenerateDSN(configs.CF.DatabaseConfig, "Main")
		db, err := DatabaseConnections(dsnMain)
		if err == nil {
			c.Locals("main", db)
			return db
		} else {
			log.Fatalf("failed to connect to Main")
		}
	}
	dbConnect, ok := dbValue.(*gorm.DB)
	if !ok {
		// หากไม่สามารถแปลงค่าให้เป็น *gorm.DB ได้ให้ทำการจัดการข้อผิดพลาดตามที่คุณต้องการ
		log.Fatalf("Error: failed to convert db1 to *gorm.DB")
	}
	fmt.Println("Connect OK")
	return dbConnect
}

func GetMainDatabase33(c *fiber.Ctx) *gorm.DB {
	dbValue := c.Locals("main")
	if dbValue == nil {
		log.Fatalf("failed to connect to Main")
	}
	dbConnect, ok := dbValue.(*gorm.DB)
	if !ok {
		// หากไม่สามารถแปลงค่าให้เป็น *gorm.DB ได้ให้ทำการจัดการข้อผิดพลาดตามที่คุณต้องการ
		log.Fatalf("Error: failed to convert db1 to *gorm.DB")
	}
	return dbConnect
}

func GetMainDatabase2(c *fiber.Ctx) *gorm.DB {
	dbValue := c.Locals("main")
	if dbValue == nil {
		dsnMain := GenerateDSN(configs.CF.DatabaseConfig, "Main")
		db, err := DatabaseConnections(dsnMain)
		if err == nil {
			c.Locals("main", db)
			return db
		} else {
			log.Fatalf("failed to connect to Main")
		}
	}
	dbConnect, ok := dbValue.(*gorm.DB)
	if !ok {
		// หากไม่สามารถแปลงค่าให้เป็น *gorm.DB ได้ให้ทำการจัดการข้อผิดพลาดตามที่คุณต้องการ
		log.Fatalf("Error: failed to convert db1 to *gorm.DB")
	}
	return dbConnect
}

func CheckDatabaseConnection() error {
	dsnMain := GenerateDSN(configs.CF.DatabaseConfig, "Main")
	db, err := DatabaseConnections(dsnMain)
	if err != nil {
		log.Printf("Failed to connect to the database: %s", err)
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// รับ *sql.DB และ error จาก db.DB()
	sqlDB, err := db.DB()
	if err != nil {
		// จัดการ error ที่ส่งกลับจาก db.DB()
		return err
	}

	// บันทึกเวลาเริ่มต้นการเชื่อมต่อ
	startTime := time.Now()

	// Ping ฐานข้อมูลด้วย *sql.DB
	err = sqlDB.Ping()
	if err != nil {
		// จัดการ error ที่ส่งกลับจาก Ping()
		return err
	}

	// คำนวณเวลาที่ใช้ในการเชื่อมต่อ
	elapsedTime := time.Since(startTime)
	log.Printf("Database connection successful. Elapsed time: %s", elapsedTime)

	// คืนค่า nil ถ้าไม่มี error หมายถึงการเชื่อมต่อสำเร็จ
	return nil
}

func CheckDatabaseStatus() {
	for {
		fmt.Println("Check Database Status")

		err := CheckDatabaseConnection()
		if err != nil {
			log.Println("Failed to connect to the database:", err)
			fmt.Println("=============================================")
			// sentry.CaptureException(err)
		} else {

			// rand.Seed(time.Now().UnixNano())
			randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
			connectTime := time.Duration(rand.Intn(5)+1) * time.Second
			time.Sleep(connectTime)

			if randGen.Intn(2) == 0 {
				// สำเร็จ
				fmt.Println("Status : OK")
				fmt.Println("=============================================")
			} else {
				// ไม่สำเร็จ
				fmt.Println("Status : SLOW")
				fmt.Println("=============================================")
			}

		}
		time.Sleep(10 * time.Second)
	}
}
