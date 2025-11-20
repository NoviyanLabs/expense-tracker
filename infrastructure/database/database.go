package database

import (
	"expense-tracker/infrastructure/loader"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *loader.Config) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second,
            LogLevel:      logger.Info,
            Colorful:      true,
        },
	)

	database := BuildPostgresDSN(cfg)
	db, err := gorm.Open(postgres.Open(database), &gorm.Config{
		Logger : newLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect database ", err)
	}

	DB = db
	return db	
}

func BuildPostgresDSN(cfg *loader.Config) string {
	db := cfg.Database
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host,
		db.Port,
		db.Username,
		db.Password,
		db.DatabaseName,
		db.SSLMode,
	)
}