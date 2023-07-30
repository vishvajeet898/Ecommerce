package api

import (
	"Ecommerce/middleware/jwt"
	"Ecommerce/order_service/models"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type OrderEndpoints struct {
	CreateOrder          endpoint.Endpoint
	GetOrderByOrderID    endpoint.Endpoint
	GetAllOrderByUserID  endpoint.Endpoint
	CancelOrderByOrderID endpoint.Endpoint
}

func MakeOrderEndpoints(svc OrderService) OrderEndpoints {
	return OrderEndpoints{
		CreateOrder:          makeCreateOrderEndpoint(svc),
		GetOrderByOrderID:    makeGetOrderByOrderIDEndpoint(svc),
		GetAllOrderByUserID:  makeGetAllOrderByUserIDEndpoint(svc),
		CancelOrderByOrderID: makeCancelOrderByOrderIDEndpoint(svc),
	}
}

func makeCreateOrderEndpoint(svc OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.CreateOrderRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.CreateOrder(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeGetOrderByOrderIDEndpoint(svc OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetOrderByOrderIDRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.GetOrderByOrderID(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeGetAllOrderByUserIDEndpoint(svc OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllOrderByUserIDRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.GetAllOrderByUserID(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeCancelOrderByOrderIDEndpoint(svc OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.CancelOrderByOrderIDRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.CancelOrderByOrderId(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func decodeCreateOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var createOrderRequest models.CreateOrderRequest
	if e := json.NewDecoder(r.Body).Decode(&createOrderRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(createOrderRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, errAuthorizationFailed
	}
	createOrderRequest.JWT = token
	return createOrderRequest, nil
}

func decodeGetOrderByOrderIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getOrderByOrderIDRequest models.GetOrderByOrderIDRequest
	if e := json.NewDecoder(r.Body).Decode(&getOrderByOrderIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(getOrderByOrderIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, errAuthorizationFailed
	}
	getOrderByOrderIDRequest.JWT = token
	return getOrderByOrderIDRequest, nil
}

func decodeGetAllOrderByUserIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getAllOrderByUserIDRequest models.GetAllOrderByUserIDRequest
	/*if e := json.NewDecoder(r.Body).Decode(&getAllOrderByUserIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(getAllOrderByUserIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}*/
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, errAuthorizationFailed
	}
	getAllOrderByUserIDRequest.JWT = token
	return getAllOrderByUserIDRequest, nil
}

func decodeCancelOrderByOrderIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var cancelOrderByOrderIDRequest models.CancelOrderByOrderIDRequest
	if e := json.NewDecoder(r.Body).Decode(&cancelOrderByOrderIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(cancelOrderByOrderIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, errAuthorizationFailed
	}
	cancelOrderByOrderIDRequest.JWT = token
	return cancelOrderByOrderIDRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewHttpService(svcEndpoints OrderEndpoints, r *mux.Router) http.Handler {

	//r := mux.NewRouter()

	r.Methods("POST").Path("/order/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.CreateOrder),
		decodeCreateOrderRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/order").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.GetOrderByOrderID),
		decodeGetOrderByOrderIDRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/orders").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.GetAllOrderByUserID),
		decodeGetAllOrderByUserIDRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/order/user/cancel").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.CancelOrderByOrderID),
		decodeCancelOrderByOrderIDRequest,
		encodeResponse,
	))

	return r
}
