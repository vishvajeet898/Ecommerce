package models

type ProductItems struct {
	ProductItemID string
	Quantity      string
	Price         string
}
type CreateOrderRequest struct {
	UserID string
	Address string
	ProductItems []ProductItems
}

type GetOrderByOrderIDRequest struct {
	OrderID string `json:"orderID"`
}

type GetAllOrderByUserIDRequest struct {
	UserID string `json:"orderID"`
}

type CancelOrderByOrderIDRequest struct {
	OrderID string `json:"orderID"`
}



