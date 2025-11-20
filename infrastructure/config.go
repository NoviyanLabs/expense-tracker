package infrastructure

import (
	"expense-tracker/infrastructure/database"
)

func Initialize() {
	database.Connect()
}