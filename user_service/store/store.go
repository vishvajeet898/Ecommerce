package store

import (
	"Ecommerce/user_service/models"
	"gorm.io/gorm"
)

type Dependency struct {
	UsersStore *EntityStore
}

type EntityStore struct {
	DB *gorm.DB
}

func NewEntityStore(DB *gorm.DB) *EntityStore {
	return &EntityStore{
		DB: DB,
	}
}

type UserStore interface {
	Create(user models.Users) error
	Update(user models.Users) error
	Delete(users models.Users) error
	GetOne(user models.Users) (*models.Users, error)
}
