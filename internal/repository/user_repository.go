package repository

import (
	"errors"
	"expense-tracker/internal/entity"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	FindAll() ([]entity.User, error)
	FindById(id uint) (*entity.User, error)
	Save(user *entity.User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]entity.User, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) FindById(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) Save(user *entity.User) error {
	return r.db.Create(user).Error
}