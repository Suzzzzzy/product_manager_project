package mapper

import (
	"example.com/m/src/domain/model"
	"time"
)

type ProductResponse struct {
	Id             int     `json:"id"`
	Category       string  `json:"category"`
	Price          float32 `json:"price"`
	Cost           float32 `json:"cost"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Barcode        string  `json:"barcode"`
	ExpirationDate string  `json:"expiration_date"`
	Size           string  `json:"size"`
	UserId         int     `json:"user_id"`
}

type ProductListResponse struct {
	Id             int       `json:"id"`
	Category       string    `json:"category"`
	Price          float32   `json:"price"`
	Cost           float32   `json:"cost"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Barcode        string    `json:"barcode"`
	ExpirationDate string    `json:"expiration_date"`
	Size           string    `json:"size"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func ToProductRes(product *model.Product) ProductResponse {
	return ProductResponse{
		Id:             product.Id,
		Category:       product.Category,
		Price:          product.Price,
		Cost:           product.Cost,
		Name:           product.Name,
		Description:    product.Description,
		Barcode:        product.Barcode,
		ExpirationDate: product.ExpirationDate.Format("2006-01-02"),
		Size:           product.Size,
		UserId:         product.UserId,
	}
}

func ToProductListRes(product []model.Product) []ProductListResponse {
	var result []ProductListResponse
	for _, prod := range product {
		result = append(result, ProductListResponse{
			Id:             prod.Id,
			Category:       prod.Category,
			Price:          prod.Price,
			Cost:           prod.Cost,
			Name:           prod.Name,
			Description:    prod.Description,
			Barcode:        prod.Barcode,
			ExpirationDate: prod.ExpirationDate.Format("2006-01-02"),
			Size:           prod.Size,
			UserId:         prod.UserId,
			CreatedAt:      prod.CreatedAt,
		})
	}
	return result
}
