package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold: time.Second,
            LogLevel:      logger.Info,
            Colorful:      true,
        },
	)

	database := "host=localhost user=admin password=admin dbname=expense_tracker port=5432"

	db, err := gorm.Open(postgres.Open(database), &gorm.Config{
		Logger : newLogger,
	})

	if err != nil {
		log.Fatal("Failed to connect database ", err)
	}

	DB = db
	
}