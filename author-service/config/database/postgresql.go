package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbGlobal *gorm.DB

func NewDB() (*gorm.DB, error) {
	if dbGlobal != nil {
		return dbGlobal, nil
	}
	var dsn string
	var dialector gorm.Dialector
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))
	dialector = postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	dbGlobal = db
	return dbGlobal, nil
}

func AutoMigrate(model ...interface{}) error {
	if dbGlobal == nil {
		return fmt.Errorf("db is not initialized")
	}
	return dbGlobal.AutoMigrate(model...)
}
