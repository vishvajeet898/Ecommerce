package store

import "Ecommerce/address_service/models"

func (e *EntityStore) Create(address models.Addresses) error {
	err := e.DB.Create(address).Error
	return err
}

func (e *EntityStore) Update(address models.Addresses) error {
	err := e.DB.Where(models.Addresses{Address_ID: address.Address_ID}).Save(address).Error
	return err
}

func (e *EntityStore) Delete(address models.Addresses) error {
	err := e.DB.Where(models.Addresses{Address_ID: address.Address_ID}).Delete(address).Error
	return err
}

func (e *EntityStore) DeleteAll(address models.Addresses) error {
	//TODO test this Deleting all query
	err := e.DB.Where(address).Delete(&models.Addresses{}).Error
	return err
}

func (e *EntityStore) GetOne(address models.Addresses) (models.Addresses, error) {
	var dbAddress models.Addresses
	err := e.DB.Where(address).First(&dbAddress).Error
	return dbAddress, err
}

func (e *EntityStore) GetAll(address models.Addresses) ([]models.Addresses, error) {
	var dbAddresses []models.Addresses
	err := e.DB.Where(address).Find(&dbAddresses).Error
	return dbAddresses, err
}
