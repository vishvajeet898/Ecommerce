package store

import "gorm.io/gorm"

type Dependency struct {
	ProductStore                   *EntityStore
	ProductItemStore               *EntityStore
	ProductVariantStore            *EntityStore
	ProductVariantValueStore       *EntityStore
	ProductVariantCombinationStore *EntityStore
}

type EntityStore struct {
	DB *gorm.DB
}

func NewEntityStore(DB *gorm.DB) *EntityStore {
	return &EntityStore{
		DB: DB,
	}
}
