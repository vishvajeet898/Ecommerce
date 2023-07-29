package store

import "Ecommerce/product_service/models"

type VariantValueStore interface {
	CreateVariantValue(models.ProductVariantValues) error
	GetOneVariantValue(models.ProductVariantValues) (models.ProductVariantValues, error)
	UpdateOneVariantValue(models.ProductVariantValues) error
	GetManyVariantValues(QueryFilter) ([]models.Product_VariantValues, error)
}

func (e *EntityStore) CreateVariantValue(productVariantValue models.ProductVariantValues) error {
	err := e.DB.Create(productVariantValue).Error
	return err
}

func (e *EntityStore) GetOneVariantValue(ProductVariantValue models.ProductVariantValues) (models.ProductVariantValues, error) {
	var dbProductVariantValue models.ProductVariantValues
	err := e.DB.Where(ProductVariantValue).First(&dbProductVariantValue).Error
	return dbProductVariantValue, err
}

func (e *EntityStore) UpdateOneVariantValue(productVariantValue models.ProductVariantValues) error {
	err := e.DB.Save(&productVariantValue).Error
	return err
}

func (e *EntityStore) GetManyVariantValues(queryFilter QueryFilter) ([]models.Product_VariantValues, error) {
	var productVariantsValues []models.Product_VariantValues
	err := e.DB.Table(queryFilter.Table).Select(queryFilter.Rows).Joins(queryFilter.Join).Where(queryFilter.Where).Find(&productVariantsValues).Error
	return productVariantsValues, err
}
