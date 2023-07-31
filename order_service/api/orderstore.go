package api

import (
	"Ecommerce/order_service/models"
	"Ecommerce/order_service/store"
	"github.com/google/uuid"
	"time"
)

type OrderStore struct {
	OrderStore     store.OrderStore
	OrderItemStore store.OrderItemStore
}

func NewOrderStoreApi(storeDependency store.Dependency) *OrderStore {
	return &OrderStore{
		OrderStore:     storeDependency.OrderStore,
		OrderItemStore: storeDependency.OrderItemStore,
	}
}

type OrderService interface {
	CreateOrder(models.CreateOrderRequest) (*models.CreateOrderResponse, error)
	GetOrderByOrderID(models.GetOrderByOrderIDRequest) (*models.GetOrderByOrderIDResponse, error)
	GetAllOrderByUserID(models.GetAllOrderByUserIDRequest) (*models.GetAllOrderByUserIDResponse, error)
	CancelOrderByOrderId(models.CancelOrderByOrderIDRequest) (*models.CancelOrderByOrderIDResponse, error)
}

func (orderstore *OrderStore) CreateOrder(createOrderRequest models.CreateOrderRequest) (*models.CreateOrderResponse, error) {
	//Create Order
	order := models.Orders{
		OrderID:   uuid.New().String(),
		UserID:    createOrderRequest.UserID,
		Timestamp: time.Now(),
		Address:   createOrderRequest.Address,
		Status:    "Packaging",
	}

	if err := orderstore.OrderStore.Create(order); err != nil {
		return nil, errUnableToCreateOrder
	}
	//Now Insert All Order Items in Order_Items

	for _, productItem := range createOrderRequest.ProductItems {
		orderItem := models.OrderItems{
			OrderItemID:   uuid.New().String(),
			OrderID:       order.OrderID,
			ProductItemID: productItem.ProductItemID,
			Quantity:      productItem.Quantity,
			Price:         productItem.Price,
		}

		if err := orderstore.OrderItemStore.CreateItem(orderItem); err != nil {
			return nil, errUnableToCreateOrder
		}
	}

	return &models.CreateOrderResponse{
		OrderID: order.OrderID,
	}, nil
}

func (orderstore *OrderStore) GetOrderByOrderID(getOrderByOrderIDRequest models.GetOrderByOrderIDRequest) (*models.GetOrderByOrderIDResponse, error) {
	//Create Order
	order := models.Orders{
		OrderID: getOrderByOrderIDRequest.OrderID,
	}

	dbOrder, err := orderstore.OrderStore.GetOne(order)
	if err != nil {
		return nil, errOrderNotExist
	}

	return &models.GetOrderByOrderIDResponse{
		Order: dbOrder,
	}, nil
}

func (orderstore *OrderStore) GetAllOrderByUserID(getAllOrderByUserIDRequest models.GetAllOrderByUserIDRequest) (*models.GetAllOrderByUserIDResponse, error) {
	//Create Order
	order := models.Orders{
		UserID: getAllOrderByUserIDRequest.UserID,
	}

	dbOrders, err := orderstore.OrderStore.GetAll(order)
	if err != nil {
		return nil, errOrderNotExist
	}
	if len(dbOrders) == 0 {
		return nil, errOrderNotExist
	}
	return &models.GetAllOrderByUserIDResponse{
		Orders: dbOrders,
	}, nil
}

func (orderstore *OrderStore) CancelOrderByOrderId(cancelOrderByOrderIDRequest models.CancelOrderByOrderIDRequest) (*models.CancelOrderByOrderIDResponse, error) {
	order := models.Orders{
		OrderID: cancelOrderByOrderIDRequest.OrderID,
	}

	//First delete all Order Items
	orderItem := models.OrderItems{
		OrderID: cancelOrderByOrderIDRequest.OrderID,
	}
	if err := orderstore.OrderItemStore.DeleteAllItem(orderItem); err != nil {
		return nil, errUnableToDeleteOrder
	}

	err := orderstore.OrderStore.Delete(order)
	if err != nil {
		return nil, errUnableToDeleteOrder
	}
	return &models.CancelOrderByOrderIDResponse{
		Error: nil,
	}, nil
}
