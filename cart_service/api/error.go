package api

import "errors"

var (
	errInternalServerError    = errors.New("internal server error")
	errAuthenticationRequired = errors.New("authentication require")
	errJsonValidation         = errors.New("json validation failed")
	errUserNotFound           = errors.New("user not found")
	errAuthenticationFailed   = errors.New("incorrect user id or password")
	errUserAlreadyExists      = errors.New("user already exists")

	errAuthorizationFailed = errors.New("authorization token not provided")
	errProductItemNotFound = errors.New("product item not found")
	errCartItemNotFound    = errors.New("cart item not found")
	errProductNotInStock   = errors.New("product item not in stock")
	errUnableToCreateItem  = errors.New("unable to add item")
	errUnableToDeleteItem  = errors.New("unable to delete item")
	errUnableToUpdateItem  = errors.New("unable to update item")
	errItemNotFound        = errors.New("cart item not found")
	errInvalidAttempt      = errors.New("invalid attempt")
	errEmptyCart           = errors.New("cart is empty")
)

type errorer interface {
	error() error
}
