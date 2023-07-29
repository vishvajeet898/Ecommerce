package store

import (
	"Ecommerce/user_service/models"
)

func (e *EntityStore) Create(user models.Users) error {
	err := e.DB.Create(&user).Error
	return err
}

func (e *EntityStore) Update(user models.Users) error {
	err := e.DB.Save(&user).Error
	return err
}

func (e *EntityStore) Delete(user models.Users) error {
	err := e.DB.Delete(user).Error
	return err
}

func (e *EntityStore) GetOne(user models.Users) (*models.Users, error) {
	var userModel models.Users
	err := e.DB.Where(user).First(&userModel).Error
	return &userModel, err
}
