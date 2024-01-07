package http

import (
	"example.com/m/src/adapter/mapper"
	"example.com/m/src/domain/model"
	"example.com/m/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
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
		router.GET("/:id", handler.GetProduct)
	}
}

type ProductRequest struct {
	Category       string  `json:"category"`
	Price          float32 `json:"price"`
	Cost           float32 `json:"cost"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Barcode        string  `json:"barcode"`
	ExpirationDate string  `json:"expiration_date"`
	Size           string  `json:"size"`
}

func (p *ProductHandler) RegisterProduct(c *gin.Context) {
	var req ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	ctx := c.Request.Context()
	token := c.Request.Header.Get("Authorization")
	// token 에서 유저정보 추출
	userId, err := utils.GetClaimByUserId(token)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	// 요청값 유효성 확인
	paredTime, err := time.Parse("2006-01-02", req.ExpirationDate)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "유통기한 날짜 형식이 맞지 않습니다.(ex.2024-01-01)", nil)
		return
	}
	if req.Size != model.Small && req.Size != model.Large {
		JSONResponse(c, http.StatusBadRequest, "size 값이 유효하지 않습니다.(ex.small or large)", nil)
		return
	}
	product := &model.Product{
		Category: req.Category, Price: req.Price, Cost: req.Cost, Name: req.Name, Description: req.Description, Barcode: req.Barcode, ExpirationDate: paredTime, Size: req.Size, UserId: userId,
	}
	newProduct, err := p.ProductUsecase.RegisterProduct(ctx, product, userId)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	result := mapper.ToProductRes(newProduct)
	JSONResponse(c, http.StatusOK, "ok", result)
}

func (p *ProductHandler) GetProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request.Context()
	token := c.Request.Header.Get("Authorization")
	// token 에서 유저정보 추출
	userId, err := utils.GetClaimByUserId(token)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	product, err := p.ProductUsecase.GetByProductId(ctx, productId, userId)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	result := mapper.ToProductRes(product)

	JSONResponse(c, http.StatusOK, "ok", result)
}
