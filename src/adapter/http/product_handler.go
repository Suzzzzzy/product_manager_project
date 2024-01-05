package http

import (
	"example.com/m/src/domain/model"
	"example.com/m/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	ProductUsecase model.ProductUsecase
}

func NewProductHandler(r *gin.Engine, u model.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: u,
	}
	router := r.Group("/products")
	{
		router.POST("", handler.RegisterProduct)

	}
}

type productRequest struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

func (u *UserHandler) RegisterProduct(c *gin.Context) {
	var req productRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c.Request.Context()
	userId, err := utils.GetClaimByUserId(ctx)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	err := u.ProductUsecase.RegisterProduct(ctx, req.Password, req.PhoneNumber)
	if err != nil {
		c.JSON(GetStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Registration successful")
}