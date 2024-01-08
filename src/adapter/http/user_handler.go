package http

import (
	"example.com/m/src/domain"
	"example.com/m/src/domain/model"
	"example.com/m/src/utils"
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
		router.GET("/signout", handler.SignOut)
	}
}

type authRequest struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (u *UserHandler) SignUp(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	ctx := c.Request.Context()
	if !utils.IsValidPhoneNumber(req.Password) {
		JSONResponse(c, http.StatusBadRequest, domain.ErrBadPhoneNumber.Error(), nil)
		return
	}
	err := u.UserUsecase.SignUp(ctx, req.Password, req.PhoneNumber)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	JSONResponse(c, http.StatusOK, "회원가입에 성공했습니다.", nil)
}

func (u *UserHandler) SignIn(c *gin.Context) {
	var req authRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	ctx := c.Request.Context()
	accessToken, err := u.UserUsecase.SignIn(ctx, req.Password, req.PhoneNumber)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	JSONResponse(c, http.StatusOK, "로그인에 성공했습니다.", gin.H{"access token": accessToken})
}

func (U *UserHandler) SignOut(c *gin.Context) {
	token, err := c.Cookie("access-token")
	if err != nil || token == "" {
		JSONResponse(c, http.StatusUnauthorized, "로그인 상태가 아닙니다.", nil)
		return
	}
	c.SetCookie("access-token", "", -1, "/", "", false, true)
	JSONResponse(c, http.StatusOK, "로그아웃 되었습니다.", nil)
}