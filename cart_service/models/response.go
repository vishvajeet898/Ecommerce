package models

type AddCartItemResponse struct {
	CartItemID string `json:"cartItemID"`
	Error      error  `json:"error"`
}

type RemoveCartItemResponse struct {
	Error error `json:"error"`
}

type cartItem struct {
	CartItemID    string `json:"cartItemID"`
	ProductItemID string `json:"productItemID"`
	Quantity      string `json:"quantity"`
	UserID        string `json:"userID"`
}
type GetAllCartItemResponse struct {
	CartItems []CartItem_ProductItem `json:"cartItems"`
}

type UpdateCartItemResponse struct {
	Error error `json:"error"`
}

type PlaceOrderResponse struct {
	OrderID string `json:"orderID"`
	Error   error  `json:"error"`
}
