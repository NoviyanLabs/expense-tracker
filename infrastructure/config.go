package infrastructure

import (
	"expense-tracker/infrastructure/database"
	"expense-tracker/infrastructure/loader"

	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, string)  {
	cfg := loader.Load()
	db := database.Connect(cfg)
	port := loader.GetServerPort(cfg)

	return db, port
}