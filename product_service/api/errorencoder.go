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
	case errUnableToAddProduct:
		w.WriteHeader(http.StatusConflict)
	case errUnableToAddProductItem:
		w.WriteHeader(http.StatusConflict)
	case errUnableToCreateVariant:
		w.WriteHeader(http.StatusConflict)
	case errUnableToCreateVariantValue:
		w.WriteHeader(http.StatusConflict)
	case errProductNotFound:
		w.WriteHeader(http.StatusNotFound)
	case errUnableToUpdate:
		w.WriteHeader(http.StatusConflict)
	case errNoAuthorizationToken:
		w.WriteHeader(http.StatusUnauthorized)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
