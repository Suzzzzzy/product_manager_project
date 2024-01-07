package mapper

import (
	"example.com/m/src/domain/model"
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
