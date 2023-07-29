package store

import (
	"Ecommerce/address_service/models"
	"gorm.io/gorm"
)

type Dependency struct {
	AddressStore *EntityStore
}

type EntityStore struct {
	DB *gorm.DB
}

func NewEntityStore(DB *gorm.DB) *EntityStore {
	return &EntityStore{
		DB: DB,
	}
}

type AddressStore interface {
	Create(addresses models.Addresses) error
	Update(models.Addresses) error
	Delete(models.Addresses) error
	GetOne(models.Addresses) (models.Addresses, error)
	GetAll(models.Addresses) ([]models.Addresses, error)
	DeleteAll(models.Addresses) error
}
