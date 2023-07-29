package api

import (
	"Ecommerce/address_service/models"
	"Ecommerce/address_service/store"
	"github.com/google/uuid"
)

type AddressStore struct {
	AddressStore store.AddressStore
}

func NewAddressStoreApi(storeDependency store.Dependency) *AddressStore {
	return &AddressStore{
		AddressStore: storeDependency.AddressStore,
	}
}

type AddressService interface {
	//JWT USER
	CreateAddress(request models.AddAddressRequest) (*models.AddAddressResponse, error)
	UpdateAddress(request models.UpdateAddressRequest) (*models.UpdateAddressResponse, error)
	DeleteAddress(request models.DeleteAddressRequest) (*models.DeleteAddressResponse, error)
	GetAllAddressByUserID(id models.GetAllAddressByUserIDRequest) (*models.GetAllAddressByUserIDResponse, error)
}

func (addressStore *AddressStore) CreateAddress(addAddressRequest models.AddAddressRequest) (*models.AddAddressResponse, error) {
	address := models.Addresses{
		User_ID:        addAddressRequest.UserID,
		Address_ID:     uuid.New().String(),
		Address_line_1: addAddressRequest.AddressLine1,
		City:           addAddressRequest.City,
		State:          addAddressRequest.State,
		Is_default:     addAddressRequest.IsDefault,
	}

	if err := addressStore.AddressStore.Create(address); err != nil {
		//TODO Err
		return nil, err
	}

	return &models.AddAddressResponse{
		Ok: nil,
	}, nil
}

func (addressStore *AddressStore) UpdateAddress(updateAddressRequest models.UpdateAddressRequest) (*models.UpdateAddressResponse, error) {
	dbAddress := models.Addresses{
		User_ID:    updateAddressRequest.UserID,
		Address_ID: updateAddressRequest.AddressID,
	}

	if _, err := addressStore.AddressStore.GetOne(dbAddress); err != nil {
		//TODO ErrAddressDoesnotExist
		return nil, err
	}

	address := models.Addresses{
		User_ID:        updateAddressRequest.UserID,
		Address_ID:     updateAddressRequest.AddressID,
		Address_line_1: updateAddressRequest.AddressLine1,
		City:           updateAddressRequest.City,
		State:          updateAddressRequest.State,
		Is_default:     updateAddressRequest.IsDefault,
	}

	if err := addressStore.AddressStore.Update(address); err != nil {
		return nil, err
	}

	return &models.UpdateAddressResponse{
		Ok: nil,
	}, nil
}

func (addressStore *AddressStore) DeleteAddress(deleteAddressRequest models.DeleteAddressRequest) (*models.DeleteAddressResponse, error) {
	//check if address exit or not
	address := models.Addresses{
		User_ID:    deleteAddressRequest.UserID,
		Address_ID: deleteAddressRequest.AddressID,
	}

	if _, err := addressStore.AddressStore.GetOne(address); err != nil {
		//TODO ErrAddressDoesnotExist
		return nil, err
	}

	//now delete it

	if err := addressStore.AddressStore.Delete(address); err != nil {
		return nil, err
	}
	return &models.DeleteAddressResponse{
		Ok: nil,
	}, nil
}

func (addressStore *AddressStore) GetAllAddressByUserID(getAllAddressByUserIDRequest models.GetAllAddressByUserIDRequest) (*models.GetAllAddressByUserIDResponse, error) {
	address := models.Addresses{
		User_ID: getAllAddressByUserIDRequest.UserID,
	}

	dbAddressItems, err := addressStore.AddressStore.GetAll(address)
	if err != nil {
		//TODO ErrAddressDoesnotExist
		return nil, err
	}

	return &models.GetAllAddressByUserIDResponse{
		Addresses: dbAddressItems,
	}, nil
}
