package store

import "Ecommerce/cart_service/models"

func (e *EntityStore) Create(cartItem models.CartItems) error {
	err := e.DB.Create(cartItem).Error
	return err
}

func (e *EntityStore) Update(cartItem models.CartItems) error {
	err := e.DB.Where(models.CartItems{CartItemID: cartItem.CartItemID}).First(&models.CartItems{}).Save(cartItem).Error
	return err
}

func (e *EntityStore) Delete(cartItem models.CartItems) error {
	err := e.DB.Where(cartItem).Delete(&models.CartItems{}).Error
	return err
}

func (e *EntityStore) DeleteAll(cartItem models.CartItems) error {
	//TODO test this Deleting all query
	err := e.DB.Where(cartItem).Delete(&models.CartItems{}).Error
	return err
}

func (e *EntityStore) GetOne(cartItem models.CartItems) (models.CartItems, error) {
	var dbCartItem models.CartItems
	err := e.DB.Where(cartItem).First(&dbCartItem).Error
	return dbCartItem, err
}

func (e *EntityStore) GetAll(queryFilter QueryFilter) ([]models.CartItem_ProductItem, error) {
	var cartItems []models.CartItem_ProductItem

	if queryFilter.Join != "" {
		err := e.DB.Table(queryFilter.Table).Select(queryFilter.Rows).Joins(queryFilter.Join).Where(queryFilter.Where).Find(&cartItems).Error
		return cartItems, err
	}
	err := e.DB.Table(queryFilter.Table).Select(queryFilter.Rows).Where(queryFilter.Where).Find(&cartItems).Error

	return cartItems, err
}
