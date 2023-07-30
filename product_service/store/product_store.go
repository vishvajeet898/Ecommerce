package store

import (
	"Ecommerce/product_service/models"
)

type ProductStore interface {
	Create(product models.Products) error
	GetOne(product models.Products) (models.Products, error)
	UpdateOne(product models.Products) error
	GetMany(product models.Products) ([]models.Products, error)
}

func (e *EntityStore) Create(product models.Products) error {
	err := e.DB.Create(product).Error
	return err
}

func (e *EntityStore) GetOne(product models.Products) (models.Products, error) {
	var dbProduct models.Products
	err := e.DB.Where(product).First(&dbProduct).Error
	return dbProduct, err
}

func (e *EntityStore) UpdateOne(product models.Products) error {
	err := e.DB.Save(&product).Error
	return err
}

func (e *EntityStore) GetMany(product models.Products) ([]models.Products, error) {
	var dbProducts []models.Products
	err := e.DB.Find(&dbProducts).Error
	return dbProducts, err
}
