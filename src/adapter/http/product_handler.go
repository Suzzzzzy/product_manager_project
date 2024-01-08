package http

import (
	"example.com/m/src/adapter/mapper"
	"example.com/m/src/domain"
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
		router.PUT("/:id", handler.UpdateProduct)
		router.DELETE("/:id", handler.DeleteProduct)
		router.GET("", handler.GetProductList)
		router.GET("/search", handler.FindProductByName)
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
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		JSONResponse(c, http.StatusUnauthorized, domain.ErrRequiredAccessToken.Error(), nil)
	}
	token := cookie.Value
	// token 에서 유저정보 추출
	userId, err := utils.VerifyToken(token)
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
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		JSONResponse(c, http.StatusUnauthorized, domain.ErrRequiredAccessToken.Error(), nil)
	}
	token := cookie.Value	// token 에서 유저정보 추출
	userId, err := utils.VerifyToken(token)
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

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		JSONResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	productId, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request.Context()
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		JSONResponse(c, http.StatusUnauthorized, domain.ErrRequiredAccessToken.Error(), nil)
	}
	token := cookie.Value	// token 에서 유저정보 추출
	userId, err := utils.VerifyToken(token)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	if _, ok := updatedFields["expiration_date"]; ok {
		expirationDate := updatedFields["expiration_date"].(string)
		_, err := time.Parse("2006-01-02", expirationDate)
		if err != nil {
			JSONResponse(c, http.StatusBadRequest, "유통기한 날짜 형식이 맞지 않습니다.(ex.2024-01-01)", nil)
			return
		}
	}
	if _, ok := updatedFields["size"]; ok {
		size := updatedFields["size"].(string)
		if size != model.Small && size != model.Large {
			JSONResponse(c, http.StatusBadRequest, "size 값이 유효하지 않습니다.(ex.small or large)", nil)
			return
		}
	}

	updatedProduct, err := p.ProductUsecase.UpdateProduct(ctx, productId, userId, updatedFields)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	result := mapper.ToProductRes(updatedProduct)
	JSONResponse(c, http.StatusOK, "ok", result)

}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("id"))

	ctx := c.Request.Context()
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		JSONResponse(c, http.StatusUnauthorized, domain.ErrRequiredAccessToken.Error(), nil)
	}
	token := cookie.Value	// token 에서 유저정보 추출
	userId, err := utils.VerifyToken(token)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	err = p.ProductUsecase.DeleteProduct(ctx, productId, userId)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}

	JSONResponse(c, http.StatusOK, "ok", nil)
}

func (p *ProductHandler) GetProductList(c *gin.Context) {
	type PageInfo struct {
		CurrentPage int `json:"current_page"`
		TotalPage   int `json:"total_page"`
	}

	ctx := c.Request.Context()
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		JSONResponse(c, http.StatusUnauthorized, domain.ErrRequiredAccessToken.Error(), nil)
	}
	token := cookie.Value	// token 에서 유저정보 추출
	userId, err := utils.VerifyToken(token)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	var page int
	if pageParam := c.Query("page"); pageParam != "" {
		parsedPage, err := strconv.Atoi(pageParam)
		if err != nil || parsedPage < 1 {
			JSONResponse(c, http.StatusBadRequest, "유효하지 않은 페이지 번호 입니다.", nil)
			return
		}
		page = parsedPage
	}

	productList, totalPage, err := p.ProductUsecase.GetProductList(ctx, userId, page)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}

	pageInfo := PageInfo{
		CurrentPage: page,
		TotalPage:   totalPage,
	}
	result := mapper.ToProductListRes(productList)

	JSONResponse(c, http.StatusOK, "ok", struct {
		Products []mapper.ProductListResponse `json:"products"`
		PageInfo PageInfo                     `json:"page_info"`
	}{
		Products: result,
		PageInfo: pageInfo,
	})
}


func (p *ProductHandler) FindProductByName(c *gin.Context) {
	ctx := c.Request.Context()
	cookie, err := c.Request.Cookie("access-token")
	if err != nil {
		JSONResponse(c, http.StatusUnauthorized, domain.ErrRequiredAccessToken.Error(), nil)
	}
	token := cookie.Value	// token 에서 유저정보 추출
	userId, err := utils.VerifyToken(token)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}

	keyword := c.Query("name")
	if keyword == ""{
		JSONResponse(c, http.StatusBadRequest, domain.ErrBadKeywordInput.Error(), nil)
		return
	}
	productList, err := p.ProductUsecase.FindProductByName(ctx, userId, keyword)
	if err != nil {
		JSONResponse(c, GetStatusCode(err), err.Error(), nil)
		return
	}
	result := mapper.ToProductListRes(productList)
	JSONResponse(c, http.StatusOK, "ok", struct {
		Products []mapper.ProductListResponse `json:"products"`
	}{
		Products: result,
	})

}