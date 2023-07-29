package api

import "errors"

var (
	errInternalServerError    = errors.New("internal server error")
	errAuthenticationRequired = errors.New("authentication require")
	errJsonValidation         = errors.New("json validation failed")
	errUserNotFound           = errors.New("user not found")
	errAuthenticationFailed   = errors.New("incorrect user id or password")
	errUserAlreadyExists      = errors.New("user already exists")
)

type errorer interface {
	error() error
}
