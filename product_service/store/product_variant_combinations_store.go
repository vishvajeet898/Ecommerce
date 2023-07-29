package store

import "Ecommerce/product_service/models"

type ProductVariantCombinationStore interface {
	CreateCombination(product models.ProductVariantCombinations) error
	GetOneCombination(product models.ProductVariantCombinations) (models.ProductVariantCombinations, error)
	UpdateOneCombination(product models.ProductVariantCombinations) error
	//GetMany(models.Products, QueryFilter) ([]interface{}, error)
}

func (e *EntityStore) CreateCombination(product models.ProductVariantCombinations) error {
	err := e.DB.Create(product).Error
	return err
}

func (e *EntityStore) GetOneCombination(product models.ProductVariantCombinations) (models.ProductVariantCombinations, error) {
	var dbProduct models.ProductVariantCombinations
	err := e.DB.Where(product).First(&dbProduct).Error
	return dbProduct, err
}

func (e *EntityStore) UpdateOneCombination(product models.ProductVariantCombinations) error {
	err := e.DB.Save(&product).Error
	return err
}
