package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errJsonValidation:
		w.WriteHeader(http.StatusBadRequest)
	case errInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
	case errAuthenticationFailed:
		w.WriteHeader(http.StatusUnauthorized)
	case errAuthenticationRequired:
		w.WriteHeader(http.StatusUnauthorized)
	case errUserNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errUserAlreadyExists:
		w.WriteHeader(http.StatusConflict)

	case errAuthorizationFailed:
		w.WriteHeader(http.StatusUnauthorized)
	case errProductItemNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errCartItemNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errProductNotInStock:
		w.WriteHeader(http.StatusConflict)
	case errEmptyCart:
		w.WriteHeader(http.StatusConflict)
	case errUnableToCreateItem:
		w.WriteHeader(http.StatusConflict)
	case errUnableToDeleteItem:
		w.WriteHeader(http.StatusConflict)
	case errUnableToUpdateItem:
		w.WriteHeader(http.StatusConflict)
	case errItemNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errInvalidAttempt:
		w.WriteHeader(http.StatusUnauthorized)

	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
