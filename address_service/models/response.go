package models

type AddAddressResponse struct {
	Ok error `json:"ok"`
}

type UpdateAddressResponse struct {
	Ok error `json:"ok"`
}

type DeleteAddressResponse struct {
	Ok error `json:"ok"`
}

type GetAllAddressByUserIDResponse struct {
	Addresses []Addresses `json:"addresses"`
}
