package store

import (
	"Ecommerce/cart_service/models"
	"gorm.io/gorm"
)

type Dependency struct {
	CartItemStore *EntityStore
}

type EntityStore struct {
	DB *gorm.DB
}

func NewEntityStore(DB *gorm.DB) *EntityStore {
	return &EntityStore{
		DB: DB,
	}
}

type CartItemStore interface {
	Create(models.CartItems) error
	Update(models.CartItems) error
	Delete(models.CartItems) error
	GetOne(models.CartItems) (models.CartItems, error)
	GetAll(models.CartItems) ([]models.CartItems, error)
	DeleteAll(models.CartItems) error
}
