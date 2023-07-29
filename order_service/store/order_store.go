package store

import "Ecommerce/order_service/models"

type OrderStore interface {
	Create(models.Orders) error
	GetOne(models.Orders) (models.Orders, error)
	UpdateOne(models.Orders) error
	GetAll(order models.Orders) ([]models.Orders, error)
	Delete(models.Orders) error
}

func (e *EntityStore) Create(order models.Orders) error {
	err := e.DB.Create(order).Error
	return err
}

func (e *EntityStore) GetOne(order models.Orders) (models.Orders, error) {
	var dbOrder models.Orders
	err := e.DB.Where(order).First(&dbOrder).Error
	return dbOrder, err
}

func (e *EntityStore) UpdateOne(order models.Orders) error {
	err := e.DB.Save(&order).Error
	return err
}

func (e *EntityStore) GetAll(order models.Orders) ([]models.Orders, error) {
	var dbOrders []models.Orders
	err := e.DB.Where(order).Find(dbOrders).Error
	return dbOrders, err
}

func (e *EntityStore) Delete(order models.Orders) error {
	err := e.DB.Delete(&order).Error
	return err
}
