package api

import (
	"errors"
)

var (
	errInternalServerError    = errors.New("internal server error")
	errAuthenticationRequired = errors.New("authentication require")
	errJsonValidation         = errors.New("json validation failed")
	errUserNotFound           = errors.New("user not found")
	errAuthenticationFailed   = errors.New("incorrect user id or password")
	errNoAuthorizationToken   = errors.New("no Authorization provided")

	errUnableToAddProduct         = errors.New("unable to add product")
	errUnableToAddProductItem     = errors.New("unable to add product item")
	errUnableToCreateVariant      = errors.New("unable to create Variant")
	errUnableToCreateVariantValue = errors.New("unable to create Variant Value")
	errProductNotFound            = errors.New("product with this ID not found")
	errUnableToUpdate             = errors.New("unable to update item")
)

type errorer interface {
	error() error
}
