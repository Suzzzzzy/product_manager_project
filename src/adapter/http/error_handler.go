package http

import (
	"example.com/m/src/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
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
	case domain.ErrUserNotFound, domain.ErrProductNotFound:
		return http.StatusNotFound
	case domain.ErrConflict, domain.ErrUserConflict:
		return http.StatusConflict
	case domain.ErrWrongPassword, domain.ErrInvalidAccessToken, domain.ErrRequiredAccessToken:
		return http.StatusUnauthorized
	case domain.ErrInvalidUser:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

func JSONResponse(c *gin.Context, code int, message string, data interface{}) {
	response := &Response{
		Meta: Meta{
			Code:    code,
			Message: message,
		},
		Data: data,
	}

	c.JSON(response.Meta.Code, response)
}
