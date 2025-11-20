package handler

import (
	"expense-tracker/internal/entity"
	"expense-tracker/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{service: service}
}


func (h *Handler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error " : "invalid id"})
		return
	}

	user, err := h.service.GetById(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error " : "Data not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) Create(c *gin.Context) {
	var req entity.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error " : err.Error()})
		return
	}

	err := h.service.Create(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error " : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

