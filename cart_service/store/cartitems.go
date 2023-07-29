package store

import "Ecommerce/cart_service/models"

func (e *EntityStore) Create(cartItem models.CartItems) error {
	err := e.DB.Create(cartItem).Error
	return err
}

func (e *EntityStore) Update(cartItem models.CartItems) error {
	err := e.DB.Save(cartItem).Error
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

func (e *EntityStore) GetAll(cartItems models.CartItems) ([]models.CartItems, error) {
	var dbCartItems []models.CartItems
	err := e.DB.Where(cartItems).Find(&dbCartItems).Error
	return dbCartItems, err
}
