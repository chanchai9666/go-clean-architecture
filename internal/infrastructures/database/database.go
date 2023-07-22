package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"eql/configs"
)

func NewDBConnection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	// ตั้งค่า connection pool
	sqlDB.SetMaxIdleConns(10)   // จำนวน connection ที่ยังไม่ถูกใช้งานสูงสุด
	sqlDB.SetMaxOpenConns(100)  // จำนวน connection ที่สามารถใช้งานพร้อมกันได้สูงสุด
	sqlDB.SetConnMaxLifetime(0) // อายุของ connection ถ้าเกินเวลาที่กำหนดจะถูกปิดและสร้างใหม่

	return db, nil
}

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
		return ""
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
