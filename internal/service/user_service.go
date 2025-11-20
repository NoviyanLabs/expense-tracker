package service

import (
	"expense-tracker/internal/entity"
	"expense-tracker/internal/repository"
)

type Service interface {
	GetAll() ([]entity.User, error)
	GetById(id uint) (*entity.User, error)
	Create(u *entity.User) error
}

type service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repository: repo}
}

// Create implements Service.
func (s *service) Create(u *entity.User) error {
	return s.repository.Save(u)
}

// GetAll implements Service.
func (s *service) GetAll() ([]entity.User, error) {
	return s.repository.FindAll()
}

// GetById implements Service.
func (s *service) GetById(id uint) (*entity.User, error) {
	return s.repository.FindById(id)
}
