package models

type AddCartItemRequest struct {
	JWT           string `json:"JWT,omitempty"`
	UserID        string `json:"UserID"`
	ProductItemID string `json:"ProductItemID"`
	Quantity      string `json:"Quantity"`
}

type RemoveCartItemRequest struct {
	JWT        string `json:"JWT,omitempty"`
	UserID     string `json:"UserID"`
	CartItemID string `json:"CartItemID"`
}

type GetAllCartItemRequest struct {
	JWT    string `json:"JWT,omitempty"`
	UserID string `json:"UserID"`
}

type UpdateCartItemRequest struct {
	JWT           string `json:"JWT,omitempty"`
	UserID        string `json:"UserID,omitempty"`
	CartItemID    string `json:"CartItemID"`
	ProductItemID string `json:"ProductItemID,omitempty"`
	Quantity      string `json:"Quantity"`
}

type ProductItems struct {
	/*	JWT           string `json:"JWT,omitempty"`
		UserID        string `json:"UserID"`*/
	ProductItemID string `json:"ProductItemID"`
	Quantity      string `json:"Quantity"`
	Price         string `json:"Price"`
}

type PlaceOrderRequest struct {
	JWT     string `json:"JWT,omitempty"`
	UserID  string `json:"UserID,omitempty"`
	Address string `json:"Address"`
	//ProductItems []ProductItems `json:"ProductItems"`
}
