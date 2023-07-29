package store

import "Ecommerce/product_service/models"

type VariantStore interface {
	CreateVariant(product models.ProductVariants) error
	GetOneVariant(product models.ProductVariants) (models.ProductVariants, error)
	UpdateOneVariant(products models.ProductVariants) error
	GetManyVariants(QueryFilter) ([]models.Product_Variants, error)
}

func (e *EntityStore) CreateVariant(productVariants models.ProductVariants) error {
	err := e.DB.Create(productVariants).Error
	return err
}

func (e *EntityStore) GetOneVariant(productVariant models.ProductVariants) (models.ProductVariants, error) {
	var dbProductVariant models.ProductVariants
	err := e.DB.Where(productVariant).First(&dbProductVariant).Error
	return dbProductVariant, err
}

func (e *EntityStore) UpdateOneVariant(productVariant models.ProductVariants) error {
	err := e.DB.Save(&productVariant).Error
	return err
}

func (e *EntityStore) GetManyVariants(queryFilter QueryFilter) ([]models.Product_Variants, error) {
	var productVariantsValues []models.Product_Variants
	err := e.DB.Table(queryFilter.Table).Select(queryFilter.Rows).Joins(queryFilter.Join).Where(queryFilter.Where).Find(&productVariantsValues).Error
	return productVariantsValues, err
}
