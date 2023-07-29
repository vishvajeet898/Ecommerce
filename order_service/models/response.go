package models

type CreateOrderResponse struct {
	OrderID string
}

type GetOrderByOrderIDResponse struct {
	Order Orders `json:"order"`
}

type GetAllOrderByUserIDResponse struct {
	Orders []Orders `json:"orders"`
}

type CancelOrderByOrderIDResponse struct {
	Error error `json:"error"`
}
