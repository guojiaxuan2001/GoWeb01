package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goweb01/global"
	"goweb01/models"
	"log"
	"time"
)

func InitDB() {
	dsn := AppConfig.Database.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
		return
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := db.AutoMigrate(&models.ExchangeRate{}); err != nil {
		log.Fatalf("Failed to migrate ExchangeRate model: %v", err)
	}

	global.Db = db
}
