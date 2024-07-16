package database

import (
	"auth-service/helpers"
	"auth-service/web/models"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitioalizeDB() *gorm.DB {

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	helpers.FatalError("Databse connect failed", err)

	sqlDB, err := db.DB()
	helpers.FatalError("Databse connect failed", err)

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func DBMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})

	return err
}
