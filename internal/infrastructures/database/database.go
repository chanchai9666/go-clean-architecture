package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"eql/configs"
)

type DBConnections struct {
	DB1 *gorm.DB //ฐานข้อมูลหลัก
	DB2 *gorm.DB //ฐานข้อมูลอื่นๆ
}

// เรียกใช้งานฐานข้อมูล
func NewDBConnections(conf configs.Config) (*DBConnections, error) {

	dsnMain1 := GenerateDSN(conf.DatabaseConfig, "Main")
	db1, err := DatabaseConnections(dsnMain1)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db1: %w", err)
	}

	dsnMain2 := GenerateDSN(conf.DatabaseConfig, "Main2")
	db2, err := DatabaseConnections(dsnMain2)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db2: %w", err)
	}

	return &DBConnections{
		DB1: db1,
		DB2: db2,
	}, nil
}

// จัดรูปแบบ DSN สำหรับใช้เชื่อมต่อฐานข้อมูล
func GenerateDSN(config configs.DatabaseConfig, dbKey string) string {
	var dbConfig struct {
		Host         string `mapstructure:"HOST"`
		Port         int    `mapstructure:"PORT"`
		Username     string `mapstructure:"USERNAME"`
		Password     string `mapstructure:"PASSWORD"`
		DatabaseName string `mapstructure:"DATABASE_NAME"`
		DriverName   string `mapstructure:"DRIVER_NAME"`
	}

	switch dbKey {
	case "Main":
		dbConfig = config.Main
	case "Main2":
		dbConfig = config.Main2
	default:
		// ถ้าไม่เจอคีย์ที่ต้องการจะใช้ในการสร้าง DSN ให้คืนค่าเป็นค่าว่าง
		return "Connection Not Found"
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DatabaseName,
	)
}

// เชื่อมต่อฐานข้อมูล
func DatabaseConnections(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// ตั้งค่า connection pool
	sqlDB.SetMaxIdleConns(10)    // จำนวน connection ที่ยังไม่ถูกใช้งานสูงสุด
	sqlDB.SetMaxOpenConns(100)   // จำนวน connection ที่สามารถใช้งานพร้อมกันได้สูงสุด
	sqlDB.SetConnMaxLifetime(10) // อายุของ connection ถ้าเกินเวลาที่กำหนดจะถูกปิดและสร้างใหม่

	return db, nil
}
