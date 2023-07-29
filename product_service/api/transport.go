package api

import (
	"Ecommerce/middleware/jwt"
	"Ecommerce/product_service/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type ProductEndpoints struct {
	CreateProduct                   endpoint.Endpoint
	CreateProductItem               endpoint.Endpoint
	GetAllProductItems              endpoint.Endpoint
	GetAllProductVariant            endpoint.Endpoint
	GetAllProductVariantByVariantID endpoint.Endpoint
	GetAllVariantValuesByProductID  endpoint.Endpoint
}

func MakeProductEndpoints(svc ProductService) ProductEndpoints {
	return ProductEndpoints{
		CreateProduct:                   makeCreateProductEndpoint(svc),
		CreateProductItem:               makeCreateProductItemEndpoint(svc),
		GetAllProductItems:              makeGetAllProductItemEndpoint(svc),
		GetAllProductVariantByVariantID: makeGetAllProductByVariantIDEndpoint(svc),
		GetAllVariantValuesByProductID:  makeGetAllVariantValueByProductIDEndpoint(svc),
	}
}

func makeCreateProductEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.AddProductRequest)
		err = svc.CreateProduct(req)
		if err != nil {
			return nil, err
		}
		return models.AddProductResponse{
			Ok: nil,
		}, err
	}
}

func makeCreateProductItemEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.AddProductItemRequest)
		err = svc.CreateProductItem(req)
		if err != nil {
			return nil, err
		}
		return models.AddProductItemResponse{
			Ok: nil,
		}, err
	}
}

func makeGetAllProductItemEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(models.GetAllProductItemsReq)

		fmt.Printf("FIne Heree")

		productItems, err := svc.GetAllProductItems()
		if err != nil {
			return nil, err
		}
		return productItems, err
	}
}

func makeGetAllProductByVariantIDEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllProductItemsByVariantIDRequest)
		productItems, err := svc.GetAllProductVariantByVariantIDs(req)
		if err != nil {
			return nil, err
		}
		return productItems, err
	}
}

func makeGetAllVariantValueByProductIDEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllVariantValueByProductIDRequest)
		productItems, err := svc.GetAllVariantValueByProductID(req)
		if err != nil {
			return nil, err
		}
		return productItems, err
	}
}

func decodeCreateProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var addProductRequest models.AddProductRequest
	if e := json.NewDecoder(r.Body).Decode(&addProductRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(addProductRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	addProductRequest.JWT = token
	return addProductRequest, nil
}

func decodeCreateProductItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var addProductItemRequest models.AddProductItemRequest
	if e := json.NewDecoder(r.Body).Decode(&addProductItemRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(addProductItemRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	addProductItemRequest.JWT = token

	return addProductItemRequest, nil
}

func decodeGetAllProductItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {

	return request, nil
}

func decodeGetAllProductByVariantIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getAllProductItemsByVariantIDRequest models.GetAllProductItemsByVariantIDRequest
	if e := json.NewDecoder(r.Body).Decode(&getAllProductItemsByVariantIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(getAllProductItemsByVariantIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	return getAllProductItemsByVariantIDRequest, nil
}

func decodeGetAllVariantByProductIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getAllVariantValueByProductIDRequest models.GetAllVariantValueByProductIDRequest
	if e := json.NewDecoder(r.Body).Decode(&getAllVariantValueByProductIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(getAllVariantValueByProductIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	return getAllVariantValueByProductIDRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func NewHttpService(svcEndpoints ProductEndpoints, r *mux.Router) http.Handler {

	//r := mux.NewRouter()
	r.Methods("POST").Path("/products/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.AdminScope})(svcEndpoints.CreateProduct),
		decodeCreateProductRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/products/item/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.AdminScope})(svcEndpoints.CreateProductItem),
		decodeCreateProductItemRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/item").Handler(httptransport.NewServer(
		svcEndpoints.GetAllProductItems,
		decodeGetAllProductItemRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/item/variants").Handler(httptransport.NewServer(
		svcEndpoints.GetAllProductVariantByVariantID,
		decodeGetAllProductByVariantIDRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/variants").Handler(httptransport.NewServer(
		svcEndpoints.GetAllVariantValuesByProductID,
		decodeGetAllVariantByProductIDRequest,
		encodeResponse,
	))
	return r
}
