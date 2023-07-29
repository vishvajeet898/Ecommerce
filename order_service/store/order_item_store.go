package store

import "Ecommerce/order_service/models"

type OrderItemStore interface {
	CreateItem(models.OrderItems) error
	GetOneItem(models.OrderItems) (models.OrderItems, error)
	UpdateOneItem(models.OrderItems) error
	GetAllItem(models.OrderItems) ([]models.OrderItems, error)
	DeleteItem(models.OrderItems) error
}

func (e *EntityStore) CreateItem(orderItem models.OrderItems) error {
	err := e.DB.Create(orderItem).Error
	return err
}

func (e *EntityStore) GetOneItem(orderItem models.OrderItems) (models.OrderItems, error) {
	var dbOrderItem models.OrderItems
	err := e.DB.Where(orderItem).First(&dbOrderItem).Error
	return dbOrderItem, err
}

func (e *EntityStore) UpdateOneItem(orderItem models.OrderItems) error {
	err := e.DB.Save(&orderItem).Error
	return err
}

func (e *EntityStore) GetAllItem(orderItem models.OrderItems) ([]models.OrderItems, error) {
	var dbOrderItems []models.OrderItems
	err := e.DB.Where(orderItem).Find(dbOrderItems).Error
	return dbOrderItems, err
}

func (e *EntityStore) DeleteItem(orderItem models.OrderItems) error {
	err := e.DB.Delete(&orderItem).Error
	return err
}
