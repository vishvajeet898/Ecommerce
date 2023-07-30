package models

type ProductItems struct {
	ProductItemID string `json:"productItemID"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
}
type CreateOrderRequest struct {
	JWT          string         `json:"JWT,omitempty"`
	UserID       string         `json:"userID,omitempty"`
	Address      string         `json:"address"`
	ProductItems []ProductItems `json:"productItems"`
}

type GetOrderByOrderIDRequest struct {
	JWT     string `json:"JWT,omitempty"`
	UserID  string `json:"userID,omitempty"`
	OrderID string `json:"orderID"`
}

type GetAllOrderByUserIDRequest struct {
	JWT    string `json:"JWT,omitempty"`
	UserID string `json:"orderID"`
}

type CancelOrderByOrderIDRequest struct {
	JWT     string `json:"JWT,omitempty"`
	UserID  string `json:"userID,omitempty"`
	OrderID string `json:"orderID"`
}
