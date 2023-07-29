package models

type AddAddressRequest struct {
	JWT          string `json:"JWT,omitempty"`
	UserID       string `json:"userID,omitempty"`
	AddressLine1 string `json:"addressLine1"`
	State        string `json:"state"`
	City         string `json:"city"`
	Country      string `json:"country"`
	IsDefault    bool   `json:"isDefault,omitempty"`
}

type UpdateAddressRequest struct {
	JWT          string `json:"JWT,omitempty"`
	UserID       string `json:"userID"`
	AddressID    string `json:"addressID"`
	AddressLine1 string `json:"addressLine1"`
	State        string `json:"state"`
	City         string `json:"city"`
	Country      string `json:"country"`
	IsDefault    bool   `json:"isDefault,omitempty"`
}

type DeleteAddressRequest struct {
	JWT       string `json:"JWT,omitempty"`
	UserID    string `json:"userID,omitempty"`
	AddressID string `json:"addressID"`
}

type GetAllAddressByUserIDRequest struct {
	JWT    string `json:"JWT,omitempty"`
	UserID string `json:"userID,omitempty"`
}
