package http

import (
	"example.com/m/src/domain"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrUserNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	case domain.ErrUserConflict:
		return http.StatusConflict
	case domain.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
