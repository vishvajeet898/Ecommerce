package api

import (
	"Ecommerce/cart_service/externals"
	"Ecommerce/cart_service/models"
	"Ecommerce/cart_service/store"
	modelOrder "Ecommerce/order_service/models"
	models2 "Ecommerce/product_service/models"
	"fmt"
	"github.com/google/uuid"
	"strconv"
)

type CartStore struct {
	CartItemStore           store.CartItemStore
	OrderExternalDependency externals.Dependency
}

func NewCartStoreApi(storeDependency store.Dependency, orderExternalDependency externals.Dependency) *CartStore {
	return &CartStore{
		CartItemStore:           storeDependency.CartItemStore,
		OrderExternalDependency: orderExternalDependency,
	}
}

type CartService interface {
	//JWT USER
	AddCartItem(models.AddCartItemRequest) (*models.AddCartItemResponse, error)
	//JWT USER
	RemoveCartItemByID(models.RemoveCartItemRequest) (*models.RemoveCartItemResponse, error)
	//JWT USER
	GetAllCartItemsByUserID(models.GetAllCartItemRequest) (*models.GetAllCartItemResponse, error)
	//JWT USER
	UpdateCartItemByID(models.UpdateCartItemRequest) (*models.UpdateCartItemResponse, error)
	//JWT USER
	//TODO empty the cart
	PlaceOrder(request models.PlaceOrderRequest) (*models.PlaceOrderResponse, error)
}

func (cartstore *CartStore) AddCartItem(addCartItemRequest models.AddCartItemRequest) (*models.AddCartItemResponse, error) {

	productItem := models2.GetProductItemByIDRequest{
		ProductItemId: addCartItemRequest.ProductItemID,
	}
	dbProductItem, err := cartstore.OrderExternalDependency.ProductService.GetProductItemByItemID(productItem)

	if err != nil {
		return nil, err
	}

	//TODO change value of quantity to INT in DB
	quantity, err := strconv.Atoi(addCartItemRequest.Quantity)
	if dbProductItem.ProductItem.Units >= quantity {
		//TODO retun ErrQuantity less
		return nil, err
	}

	cartItem := models.CartItems{
		CartItemID:    uuid.New().String(),
		UserID:        addCartItemRequest.UserID,
		ProductItemID: addCartItemRequest.ProductItemID,
		Quantity:      addCartItemRequest.Quantity,
	}

	if err := cartstore.CartItemStore.Create(cartItem); err != nil {
		return nil, err
	}

	return &models.AddCartItemResponse{
		CartItemID: cartItem.CartItemID,
		Error:      nil,
	}, nil
}

func (cartstore *CartStore) RemoveCartItemByID(removeCartItemRequest models.RemoveCartItemRequest) (*models.RemoveCartItemResponse, error) {
	cartItem := models.CartItems{
		UserID:     removeCartItemRequest.UserID,
		CartItemID: removeCartItemRequest.CartItemID,
	}
	if err := cartstore.CartItemStore.Delete(cartItem); err != nil {
		return nil, err
	}
	return &models.RemoveCartItemResponse{
		Error: nil,
	}, nil
}

func (cartstore *CartStore) GetAllCartItemsByUserID(getAllCartItemRequest models.GetAllCartItemRequest) (*models.GetAllCartItemResponse, error) {
	cartItem := models.CartItems{
		UserID: getAllCartItemRequest.UserID,
	}
	cartItems, err := cartstore.CartItemStore.GetAll(cartItem)
	if err != nil {
		return nil, err
	}

	return &models.GetAllCartItemResponse{
		CartItems: cartItems,
	}, nil
}

func (cartstore *CartStore) UpdateCartItemByID(updateCartItemRequest models.UpdateCartItemRequest) (*models.UpdateCartItemResponse, error) {
	dbCartItem, err := cartstore.CartItemStore.GetOne(models.CartItems{
		CartItemID: updateCartItemRequest.CartItemID,
	})

	//If errr then item either does not exist
	if err != nil {
		return nil, err
	}

	cartItem := models.CartItems{
		CartItemID:    updateCartItemRequest.CartItemID,
		UserID:        dbCartItem.UserID,
		Quantity:      updateCartItemRequest.Quantity,
		ProductItemID: dbCartItem.ProductItemID,
	}

	err = cartstore.CartItemStore.Update(cartItem)
	if err != nil {
		return nil, err
	}

	return &models.UpdateCartItemResponse{
		Error: nil,
	}, nil
}

func (cartstore *CartStore) PlaceOrder(placeOrderRequest models.PlaceOrderRequest) (*models.PlaceOrderResponse, error) {

	//All Cart Items
	cartItem := models.CartItems{
		UserID: placeOrderRequest.UserID,
	}
	dbCartItems, err := cartstore.CartItemStore.GetAll(cartItem)
	if err != nil {
		return nil, err
	}

	if len(dbCartItems) == 0 {
		return nil, fmt.Errorf("Cart Empty")
	}

	productItems := make([]modelOrder.ProductItems, 0)
	for _, cartItem := range dbCartItems {

		//optimization Get bulk select instead of loop
		productItem, err := cartstore.OrderExternalDependency.ProductService.GetProductItemByItemID(models2.GetProductItemByIDRequest{ProductItemId: cartItem.ProductItemID})
		if err != nil {
			return nil, err
		}

		productItems = append(productItems, modelOrder.ProductItems{
			ProductItemID: productItem.ProductItem.ProductItemId,
			Price:         productItem.ProductItem.Price,
			Quantity:      cartItem.Quantity,
		})
	}

	orderRequest := modelOrder.CreateOrderRequest{
		UserID:       placeOrderRequest.UserID,
		Address:      placeOrderRequest.Address,
		ProductItems: productItems,
	}

	placedOrder, err := cartstore.OrderExternalDependency.OrderService.CreateOrder(orderRequest)
	if err != nil {
		return nil, err
	}

	//Emptying Cart after placing order
	/*cartItem := models.CartItems{
		UserID: placeOrderRequest.UserID,
	}*/
	if err := cartstore.CartItemStore.DeleteAll(cartItem); err != nil {
		return nil, err
	}

	return &models.PlaceOrderResponse{
		OrderID: placedOrder.OrderID,
		Error:   nil,
	}, nil
}
