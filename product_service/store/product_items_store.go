package store

import (
	"Ecommerce/product_service/models"
)

type ProductItemsStore interface {
	CreateItem(product models.ProductItems) error
	GetOneItem(product models.ProductItems) (models.ProductItems, error)
	UpdateOneItem(products models.ProductItems) error
	GetAllItems(QueryFilter) ([]models.ProductItems, error)
	//GetMany(models.ProductItems, QueryFilter) ([]models.ProductItems, error)
}

func (e *EntityStore) CreateItem(productItem models.ProductItems) error {
	err := e.DB.Create(productItem).Error
	return err
}

func (e *EntityStore) GetOneItem(productItem models.ProductItems) (models.ProductItems, error) {
	var dbProduct models.ProductItems
	err := e.DB.Where(productItem).First(&dbProduct).Error
	return dbProduct, err
}

func (e *EntityStore) UpdateOneItem(productItem models.ProductItems) error {
	err := e.DB.Where(models.ProductItems{ProductItemId: productItem.ProductItemId}).Save(&productItem).Error
	return err
}

func (e *EntityStore) GetAllItems(queryFilter QueryFilter) ([]models.ProductItems, error) {
	var productItems []models.ProductItems
	if queryFilter.Where != "" {
		err := e.DB.Where(queryFilter.Where).Find(&productItems).Error
		return productItems, err
	}
	err := e.DB.Find(&productItems).Error
	return productItems, err
}
