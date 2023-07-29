package store

import "gorm.io/gorm"

type Dependency struct {
	OrderStore     *EntityStore
	OrderItemStore *EntityStore
}

type EntityStore struct {
	DB *gorm.DB
}

func NewEntityStore(DB *gorm.DB) *EntityStore {
	return &EntityStore{
		DB: DB,
	}
}
