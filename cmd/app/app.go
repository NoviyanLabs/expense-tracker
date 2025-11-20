package app

import (
	"expense-tracker/infrastructure"
	"expense-tracker/internal/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	db, port := infrastructure.Initialize()

	fmt.Print(port)
	r := gin.Default()
	router.RegisterRoute(r, db)
	r.Run(":"+port)
}