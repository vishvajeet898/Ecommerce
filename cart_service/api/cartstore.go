package api

import (
	"Ecommerce/cart_service/externals"
	"Ecommerce/cart_service/models"
	"Ecommerce/cart_service/store"
	modelOrder "Ecommerce/order_service/models"
	models2 "Ecommerce/product_service/models"
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
	AddCartItem(models.AddCartItemRequest) (*models.AddCartItemResponse, error)
	RemoveCartItemByID(models.RemoveCartItemRequest) (*models.RemoveCartItemResponse, error)
	GetAllCartItemsByUserID(models.GetAllCartItemRequest) (*models.GetAllCartItemResponse, error)
	UpdateCartItemByID(models.UpdateCartItemRequest) (*models.UpdateCartItemResponse, error)
	PlaceOrder(request models.PlaceOrderRequest) (*models.PlaceOrderResponse, error)
}

func (cartstore *CartStore) AddCartItem(addCartItemRequest models.AddCartItemRequest) (*models.AddCartItemResponse, error) {

	productItem := models2.GetProductItemByIDRequest{
		ProductItemId: addCartItemRequest.ProductItemID,
	}
	dbProductItem, err := cartstore.OrderExternalDependency.ProductService.GetProductItemByItemID(productItem)

	if err != nil {
		return nil, errProductItemNotFound
	}

	//TODO change value of quantity to INT in DB
	quantity, err := strconv.Atoi(addCartItemRequest.Quantity)
	if dbProductItem.ProductItem.Units <= quantity {
		//TODO retun ErrQuantity less
		return nil, errProductNotInStock
	}

	cartItem := models.CartItems{
		CartItemID:    uuid.New().String(),
		UserID:        addCartItemRequest.UserID,
		ProductItemID: addCartItemRequest.ProductItemID,
		Quantity:      addCartItemRequest.Quantity,
	}

	if err := cartstore.CartItemStore.Create(cartItem); err != nil {
		return nil, errUnableToCreateItem
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
	dbCartItem, err := cartstore.CartItemStore.GetOne(cartItem)
	if err != nil {
		return nil, errCartItemNotFound
	}

	if dbCartItem.UserID != removeCartItemRequest.UserID {
		return nil, errInvalidAttempt
	}

	if err := cartstore.CartItemStore.Delete(cartItem); err != nil {
		return nil, errUnableToDeleteItem
	}
	return &models.RemoveCartItemResponse{
		Error: nil,
	}, nil
}

func (cartstore *CartStore) GetAllCartItemsByUserID(getAllCartItemRequest models.GetAllCartItemRequest) (*models.GetAllCartItemResponse, error) {

	querryFiler := store.QueryFilter{
		Table: "cart_items",
		Rows:  "cart_items.cart_item_id, cart_items.product_item_id, cart_items.quantity, product_items.name, product_items.price",
		Join:  "inner join product_items on cart_items.product_item_id = product_items.product_item_id",
		Where: "cart_items.user_id = '" + getAllCartItemRequest.UserID + "';",
	}
	cartItems, err := cartstore.CartItemStore.GetAll(querryFiler)
	if err != nil {
		return nil, errInternalServerError
	}

	return &models.GetAllCartItemResponse{
		CartItems: cartItems,
	}, nil
}

func (cartstore *CartStore) UpdateCartItemByID(updateCartItemRequest models.UpdateCartItemRequest) (*models.UpdateCartItemResponse, error) {
	dbCartItem, err := cartstore.CartItemStore.GetOne(models.CartItems{
		CartItemID: updateCartItemRequest.CartItemID,
	})

	//If err then item either does not exist
	if err != nil {
		return nil, errItemNotFound
	}

	if dbCartItem.UserID != updateCartItemRequest.UserID {
		return nil, errInvalidAttempt
	}

	cartItem := models.CartItems{
		CartItemID:    updateCartItemRequest.CartItemID,
		UserID:        dbCartItem.UserID,
		Quantity:      updateCartItemRequest.Quantity,
		ProductItemID: dbCartItem.ProductItemID,
	}

	err = cartstore.CartItemStore.Update(cartItem)
	if err != nil {
		return nil, errUnableToUpdateItem
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

	queryFiler := store.QueryFilter{
		Table: "cart_items",
		Rows:  "*",
		Where: "cart_items.user_id = '" + placeOrderRequest.UserID + "';",
	}

	dbCartItems, err := cartstore.CartItemStore.GetAll(queryFiler)
	if err != nil {
		return nil, errInternalServerError
	}

	if len(dbCartItems) == 0 {
		return nil, errEmptyCart
	}

	productItems := make([]modelOrder.ProductItems, 0)
	for _, cartItem := range dbCartItems {

		//optimization Get bulk select instead of loop
		productItem, err := cartstore.OrderExternalDependency.ProductService.GetProductItemByItemID(models2.GetProductItemByIDRequest{ProductItemId: cartItem.ProductItemID})
		if err != nil {
			return nil, errProductItemNotFound
		}

		quantity, _ := strconv.Atoi(cartItem.Quantity)

		//Reducing the quantity of Product Item
		_, err = cartstore.OrderExternalDependency.ProductService.UpdateProductItem(models2.UpdateProductItemRequest{
			ProductItemId: productItem.ProductItem.ProductItemId,
			Units:         productItem.ProductItem.Units - quantity,
			Name:          productItem.ProductItem.Name,
			Price:         productItem.ProductItem.Price,
		})
		if err != nil {
			return nil, errInternalServerError
		}

		productItems = append(productItems, modelOrder.ProductItems{
			ProductItemID: productItem.ProductItem.ProductItemId,
			Price:         productItem.ProductItem.Price,
			Quantity:      cartItem.Quantity,
		})
	}

	//Creating order
	orderRequest := modelOrder.CreateOrderRequest{
		UserID:       placeOrderRequest.UserID,
		Address:      placeOrderRequest.Address,
		ProductItems: productItems,
	}

	placedOrder, err := cartstore.OrderExternalDependency.OrderService.CreateOrder(orderRequest)
	if err != nil {
		return nil, errInternalServerError
	}

	//Emptying Cart after placing order
	if err := cartstore.CartItemStore.DeleteAll(cartItem); err != nil {
		return nil, errInternalServerError
	}

	return &models.PlaceOrderResponse{
		OrderID: placedOrder.OrderID,
		Error:   nil,
	}, nil
}
