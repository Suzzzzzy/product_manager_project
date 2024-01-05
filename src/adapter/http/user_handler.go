package http

import (
	"example.com/m/src/domain/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
		router.POST("/signup", handler.SignUp)
		router.POST("/signin", handler.SignIn)
		router.POST("/signout")
	}
}

type authRequest struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (u *UserHandler) SignUp(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c.Request.Context()
	err := u.UserUsecase.SignUp(ctx, req.Password, req.PhoneNumber)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Registration successful")
}

func (u *UserHandler) SignIn(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c.Request.Context()
	accessToken, err := u.UserUsecase.SignIn(ctx, req.Password, req.PhoneNumber)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(http.StatusOK, gin.H{"access token": accessToken})
}
