package router

import (
	"expense-tracker/internal/handler"
	"expense-tracker/internal/repository"
	"expense-tracker/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoute(r *gin.Engine, db *gorm.DB) {

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	u := r.Group("/users") 
	{
		u.GET("/", handler.GetAll)
		u.GET("/:id", handler.GetById)
		u.POST("/", handler.Create)
	}

}