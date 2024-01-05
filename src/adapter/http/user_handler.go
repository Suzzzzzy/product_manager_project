package http

import (
	"example.com/m/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	UserUsecase model.UserUsecase
}

func NewUserHandler(r *gin.Engine, u model.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: u,
	}
	router := r.Group("/auth")
	{
		router.POST("/signin", handler.SignIn)
		router.POST("/signup")
		router.POST("/signout")
	}
}

type signinRequest struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (u *UserHandler) SignIn(c *gin.Context) {
	var req signinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c.Request.Context()
	err := u.UserUsecase.CreateUser(ctx, req.Password, req.PhoneNumber)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Registration successful")
}
