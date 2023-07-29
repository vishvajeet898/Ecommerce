package api

import (
	"Ecommerce/address_service/models"
	"Ecommerce/middleware/jwt"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type AddressEndpoints struct {
	CreateAddress endpoint.Endpoint
	UpdateAddress endpoint.Endpoint
	DeleteAddress endpoint.Endpoint
	GetAllAddress endpoint.Endpoint
}

func MakeAddressEndpoints(svc AddressService) AddressEndpoints {
	return AddressEndpoints{
		CreateAddress: makeCreateAddressEndpoint(svc),
		UpdateAddress: makeUpdateAddressEndpoint(svc),
		DeleteAddress: makeDeleteAddressEndpoint(svc),
		GetAllAddress: makeGetAllAddressEndpoint(svc),
	}
}

func makeCreateAddressEndpoint(svc AddressService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.AddAddressRequest)

		req.UserID = ctx.Value("userID").(string)

		response, err = svc.CreateAddress(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeUpdateAddressEndpoint(svc AddressService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.UpdateAddressRequest)

		req.UserID = ctx.Value("userID").(string)

		response, err = svc.UpdateAddress(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeDeleteAddressEndpoint(svc AddressService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.DeleteAddressRequest)

		req.UserID = ctx.Value("userID").(string)

		response, err = svc.DeleteAddress(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func makeGetAllAddressEndpoint(svc AddressService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetAllAddressByUserIDRequest)

		req.UserID = ctx.Value("userID").(string)

		response, err = svc.GetAllAddressByUserID(req)
		if err != nil {
			return nil, err
		}
		return response, err
	}
}

func decodeCreateAddressRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var addAddressRequest models.AddAddressRequest
	if e := json.NewDecoder(r.Body).Decode(&addAddressRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(addAddressRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	addAddressRequest.JWT = token

	return addAddressRequest, nil
}

func decodeUpdateAddressRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var updateAddressRequest models.UpdateAddressRequest
	if e := json.NewDecoder(r.Body).Decode(&updateAddressRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(updateAddressRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	updateAddressRequest.JWT = token
	return updateAddressRequest, nil
}

func decodeDeleteAddressRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var deleteAddressRequest models.DeleteAddressRequest
	if e := json.NewDecoder(r.Body).Decode(&deleteAddressRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(deleteAddressRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	deleteAddressRequest.JWT = token
	return deleteAddressRequest, nil
}

func decodeGetAllAddressRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var getAllAddressByUserIDRequest models.GetAllAddressByUserIDRequest
	if e := json.NewDecoder(r.Body).Decode(&getAllAddressByUserIDRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(getAllAddressByUserIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	getAllAddressByUserIDRequest.JWT = token
	return getAllAddressByUserIDRequest, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// NewHttpService HTTP Server
func NewHttpService(svcEndpoints AddressEndpoints, r *mux.Router) http.Handler {

	//r := mux.NewRouter()
	r.Methods("POST").Path("/address/add").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.CreateAddress),
		decodeCreateAddressRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/address/update").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.UpdateAddress),
		decodeUpdateAddressRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/address/delete").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.DeleteAddress),
		decodeDeleteAddressRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/address/all").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.GetAllAddress),
		decodeGetAllAddressRequest,
		encodeResponse,
	))

	return r
}
