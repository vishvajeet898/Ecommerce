package api

import (
	"Ecommerce/cart_service/models"
	"Ecommerce/middleware/jwt"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type CartEndpoints struct {
	AddCartItem            endpoint.Endpoint
	RemoveCartItem         endpoint.Endpoint
	GetAllCartItemByUserID endpoint.Endpoint
	UpdateCartItemByID     endpoint.Endpoint
	PlacerOrder            endpoint.Endpoint
}

func MakeOrderEndpoints(svc CartService) CartEndpoints {
	return CartEndpoints{
		AddCartItem:            makeAddCartItemEndpoint(svc),
		RemoveCartItem:         makeRemoveCartItemEndpoint(svc),
		GetAllCartItemByUserID: makeGetAllCartItemByUserIDEndpoint(svc),
		UpdateCartItemByID:     makeUpdateCartItemByIDEndpoint(svc),
		PlacerOrder:            makePlaceOrderEndpoint(svc),
	}
}

func makeAddCartItemEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.AddCartItemRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.AddCartItem(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeRemoveCartItemEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.RemoveCartItemRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.RemoveCartItemByID(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeGetAllCartItemByUserIDEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllCartItemRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.GetAllCartItemsByUserID(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeUpdateCartItemByIDEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.UpdateCartItemRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.UpdateCartItemByID(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makePlaceOrderEndpoint(svc CartService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.PlaceOrderRequest)
		req.UserID = ctx.Value("userID").(string)
		response, err = svc.PlaceOrder(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func decodeAddCartItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var addCartItemRequest models.AddCartItemRequest
	if e := json.NewDecoder(r.Body).Decode(&addCartItemRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(addCartItemRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errAuthorizationFailed
	}
	addCartItemRequest.JWT = token

	return addCartItemRequest, nil
}

func decodeRemoveCartItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var removeCartItemRequest models.RemoveCartItemRequest
	if e := json.NewDecoder(r.Body).Decode(&removeCartItemRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(removeCartItemRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errAuthorizationFailed
	}
	removeCartItemRequest.JWT = token

	return removeCartItemRequest, nil
}

func decodeGetAllCartItemByUserIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getAllCartItemRequest models.GetAllCartItemRequest
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errAuthorizationFailed
	}
	getAllCartItemRequest.JWT = token
	return getAllCartItemRequest, nil
}

func decodeUpdateCartItemByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var updateCartItemRequest models.UpdateCartItemRequest
	if e := json.NewDecoder(r.Body).Decode(&updateCartItemRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(updateCartItemRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errAuthorizationFailed
	}
	updateCartItemRequest.JWT = token
	return updateCartItemRequest, nil
}

func decodePlacerOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var placeOrderRequest models.PlaceOrderRequest
	if e := json.NewDecoder(r.Body).Decode(&placeOrderRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(placeOrderRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, errAuthorizationFailed
	}
	placeOrderRequest.JWT = token
	return placeOrderRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewHttpService(svcEndpoints CartEndpoints, r *mux.Router) http.Handler {

	//r := mux.NewRouter()

	r.Methods("POST").Path("/cart/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.AddCartItem),
		decodeAddCartItemRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/cart/remove").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.RemoveCartItem),
		decodeRemoveCartItemRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/cart/all").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.GetAllCartItemByUserID),
		decodeGetAllCartItemByUserIDRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/cart/update").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.UpdateCartItemByID),
		decodeUpdateCartItemByIDRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/cart/placeOrder").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.PlacerOrder),
		decodePlacerOrderRequest,
		encodeResponse,
	))

	return r
}
