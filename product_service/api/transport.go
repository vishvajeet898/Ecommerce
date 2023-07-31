package api

import (
	"Ecommerce/middleware/jwt"
	"Ecommerce/product_service/models"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type ProductEndpoints struct {
	CreateProduct                    endpoint.Endpoint
	CreateProductItem                endpoint.Endpoint
	GetAllProductItems               endpoint.Endpoint
	GetAllProductInformation         endpoint.Endpoint
	GetAllProductVariantsByVariantID endpoint.Endpoint
	GetAllVariantValuesByProductID   endpoint.Endpoint
	UpdateProductItem                endpoint.Endpoint
}

func MakeProductEndpoints(svc ProductService) ProductEndpoints {
	return ProductEndpoints{
		CreateProduct:                    makeCreateProductEndpoint(svc),
		CreateProductItem:                makeCreateProductItemEndpoint(svc),
		GetAllProductItems:               makeGetAllProductItemEndpoint(svc),
		GetAllProductInformation:         makeGetAllProductEndpoint(svc),
		GetAllProductVariantsByVariantID: makeGetAllProductVariantByVariantIDEndpoint(svc),
		GetAllVariantValuesByProductID:   makeGetAllVariantValueByProductIDEndpoint(svc),
		UpdateProductItem:                makeUpdateProductItemEndpoint(svc),
	}
}

func makeCreateProductEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.AddProductRequest)
		resp, err := svc.CreateProduct(req)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}

func makeCreateProductItemEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.AddProductItemRequest)
		resp, err := svc.CreateProductItem(req)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}

func makeGetAllProductItemEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(models.GetAllProductItemsReq)
		productItems, err := svc.GetAllProductItems()
		if err != nil {
			return nil, err
		}
		return productItems, err
	}
}

func makeGetAllProductEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//req := request.(models.GetAllProductItemsReq)
		productItems, err := svc.GetAllProduct()
		if err != nil {
			return nil, err
		}
		return productItems, err
	}
}

func makeGetAllProductVariantByVariantIDEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllProductItemsByVariantValueIDRequest)
		productItems, err := svc.GetAllProductItemsByVariantValueID(req)
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

func makeGetAllProductItemByVariantIDEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllProductItemsByVariantValueIDRequest)
		productItems, err := svc.GetAllProductItemsByVariantValueID(req)
		if err != nil {
			return nil, err
		}
		return productItems, err
	}
}

func makeUpdateProductItemEndpoint(svc ProductService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.UpdateProductItemRequest)
		ok, err := svc.UpdateProductItem(req)
		if err != nil {
			return nil, err
		}
		return ok, err
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
		return nil, errNoAuthorizationToken
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
		return nil, errNoAuthorizationToken
	}
	addProductItemRequest.JWT = token

	return addProductItemRequest, nil
}

func decodeGetAllProductItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return request, nil
}

func decodeGetAllProductRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return request, nil
}

func decodeGetAllProductVariantsByVariantIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getAllProductIVariantsByVariantIDRequest models.GetAllProductItemsByVariantValueIDRequest
	if e := json.NewDecoder(r.Body).Decode(&getAllProductIVariantsByVariantIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(getAllProductIVariantsByVariantIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	return getAllProductIVariantsByVariantIDRequest, nil
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

func decodeUpdateProductItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var updateProductItemRequest models.UpdateProductItemRequest
	if e := json.NewDecoder(r.Body).Decode(&updateProductItemRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(updateProductItemRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, errNoAuthorizationToken
	}
	updateProductItemRequest.JWT = token

	return updateProductItemRequest, nil
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
	r.Methods("POST").Path("/product/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.AdminScope})(svcEndpoints.CreateProduct),
		decodeCreateProductRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/product/item/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.AdminScope})(svcEndpoints.CreateProductItem),
		decodeCreateProductItemRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/item").Handler(httptransport.NewServer(
		svcEndpoints.GetAllProductItems,
		decodeGetAllProductItemRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/all").Handler(httptransport.NewServer(
		svcEndpoints.GetAllProductInformation,
		decodeGetAllProductRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/item/variants").Handler(httptransport.NewServer(
		svcEndpoints.GetAllProductVariantsByVariantID,
		decodeGetAllProductVariantsByVariantIDRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/product/variants").Handler(httptransport.NewServer(
		svcEndpoints.GetAllVariantValuesByProductID,
		decodeGetAllVariantByProductIDRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/product/item/update").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.AdminScope})(svcEndpoints.UpdateProductItem),
		decodeUpdateProductItemRequest,
		encodeResponse,
	))
	return r
}
