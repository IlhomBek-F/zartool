package internal

import (
	"errors"
	"net/http"
)

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
	// Error user does not exist
	ErrUserNotFound = errors.New("user does not exist")
)

func GetErrorCode(err error) (int, string) {
	switch {
	case errors.Is(err, ErrInternalServerError):
		return http.StatusInternalServerError, ErrInternalServerError.Error()
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound, ErrNotFound.Error()
	case errors.Is(err, ErrConflict):
		return http.StatusConflict, ErrConflict.Error()
	case errors.Is(err, ErrBadParamInput):
		return http.StatusBadRequest, ErrBadParamInput.Error()
	case errors.Is(err, ErrUserNotFound):
		return http.StatusNotFound, ErrUserNotFound.Error()
	default:
		return http.StatusInternalServerError, ErrInternalServerError.Error()
	}
}
