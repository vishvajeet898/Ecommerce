package api

import (
	"Ecommerce/middleware/jwt"
	"Ecommerce/user_service/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"net/http"
)

type UserEndpoints struct {
	Create         endpoint.Endpoint
	GetByID        endpoint.Endpoint
	UpdateUser     endpoint.Endpoint
	GetUserByID    endpoint.Endpoint
	DeleteUserByID endpoint.Endpoint
}

func MakeUserEndpoints(svc UsersService) UserEndpoints {

	return UserEndpoints{
		Create:         makeCreateSignUpUserEndpoint(svc),
		GetByID:        makeCreateSignInUserEndpoint(svc),
		UpdateUser:     makeUpdateUserEndpoint(svc),
		GetUserByID:    makeGetUserEndpoint(svc),
		DeleteUserByID: makeDeleteUserEndpoint(svc),
	}
}

// Endpoints
func makeCreateSignUpUserEndpoint(svc UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.SignUpUserRequest)
		jwt, err := svc.SignUpUser(req)
		if err != nil {
			return models.SignUpUserResponse{
				Error: err}, err
		}
		return models.SignUpUserResponse{
			Jwt: jwt}, err
	}
}

func makeCreateSignInUserEndpoint(svc UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.SignInUserRequest)
		jwt, err := svc.LoginUser(req)
		if err != nil {
			return models.SignInUserResponse{
				Error: err,
			}, err
		}

		return models.SignUpUserResponse{
			Jwt: jwt,
		}, err
	}
}

func makeUpdateUserEndpoint(svc UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.UpdateUserRequest)

		req.UserID = ctx.Value("userID").(string)

		err = svc.UpdateUser(req)
		if err != nil {
			return models.UpdateUserResponse{
				Ok: err,
			}, err
		}

		return models.UpdateUserResponse{
			Ok: nil,
		}, err
	}
}

func makeGetUserEndpoint(svc UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.GetUserByIDRequest)
		req.User_ID = ctx.Value("userID").(string)
		user, err := svc.GetUserById(req)
		if err != nil {
			return nil, err
		}

		return models.GetUserByIDResponse{
			User: *user,
		}, err
	}
}

func makeDeleteUserEndpoint(svc UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.DeleteUserByIDRequest)

		req.User_ID = ctx.Value("userID").(string)
		err = svc.DeleteUser(req)
		if err != nil {
			return nil, err
		}

		return models.DeleteUserByIDResponse{
			Ok: nil,
		}, err
	}
}

// EncodeDecode requests
func decodeSignUpUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var userSignUpRequest models.SignUpUserRequest
	if e := json.NewDecoder(r.Body).Decode(&userSignUpRequest); e != nil {
		return nil, e
	}

	//Validating the fields of the struct
	v := validator.New()
	err = v.Struct(userSignUpRequest)
	if err != nil {
		return nil, errJsonValidation
	}
	return userSignUpRequest, nil
}

func decodeLoginUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var userSignInRequest models.SignInUserRequest
	if e := json.NewDecoder(r.Body).Decode(&userSignInRequest); e != nil {
		return nil, e
	}
	v := validator.New()
	err = v.Struct(userSignInRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	return userSignInRequest, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var updateRequest models.UpdateUserRequest
	if e := json.NewDecoder(r.Body).Decode(&updateRequest); e != nil {
		return nil, e
	}

	v := validator.New()
	err = v.Struct(updateRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, err
	}
	updateRequest.JWT = token

	return updateRequest, nil
}

func decodeGetUserByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var userByIDRequest models.GetUserByIDRequest
	token := r.Header.Get("Authorization")
	if token == "" {
		//TODO return ERR
		return nil, fmt.Errorf("no Authorization provided")
	}
	userByIDRequest.JWT = token
	return userByIDRequest, nil
}

func decodeDeleteUserByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var deleteUserByIDRequest models.DeleteUserByIDRequest
	if e := json.NewDecoder(r.Body).Decode(&deleteUserByIDRequest); e != nil {
		return nil, e
	}
	v := validator.New()
	err = v.Struct(deleteUserByIDRequest)
	if err != nil {
		return nil, errJsonValidation
	}

	return deleteUserByIDRequest, nil
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
func NewHttpService(svcEndpoints UserEndpoints, r *mux.Router) http.Handler {

	//r := mux.NewRouter()
	r.Methods("POST").Path("/users/signup").Handler(httptransport.NewServer(
		svcEndpoints.Create,
		decodeSignUpUserRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/users/login").Handler(httptransport.NewServer(
		svcEndpoints.GetByID,
		decodeLoginUserRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/users/update").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.UpdateUser),
		decodeUpdateUserRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/users/user").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.GetUserByID),
		decodeGetUserByIDRequest,
		encodeResponse,
	))

	r.Methods("DELETE").Path("/users/delete").Handler(httptransport.NewServer(
		jwt.NewAuthMiddleware([]string{jwt.UserScope})(svcEndpoints.DeleteUserByID),
		decodeDeleteUserByIDRequest,
		encodeResponse,
	))
	return r
}
