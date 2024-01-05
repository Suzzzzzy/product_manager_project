package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	// 상품의 고유 아이디
	ID int `gorm:"primaryKey;autoIncrement;column:id;type:INT8" json:"id"`
	// 카테고리
	Category string `gorm:"column:category;type:VARCHAR(50) NOT NULL;" json:"category"`
	// 가격
	Price float32 `gorm:"column:price;type:FLOAT NOT NULL;" json:"price"`
	// 원가
	Cost float32 `gorm:"column:cost;type:FLOAT NOT NULL;" json:"cost"`
	// 이름
	Name string `gorm:"column:name;type:VARCHAR(50) NOT NULL;" json:"name"`
	// 설명
	Description string `gorm:"column:description;type:VARCHAR(255) NOT NULL;" json:"description"`
	// 바코드
	Barcode string `gorm:"column:barcode;type:VARCHAR(255) NOT NULL;" json:"barcode"`
	// 유통기한
	ExpirationDate time.Time `gorm:"column:expiration_date;type:DATETIME NOT NULL;" json:"expiration_date"`
	// 사이즈
	Size string `gorm:"column:size;type:VARCHAR(50) NOT NULL;" json:"size"`
	// 상품을 등록한 유저의 고유값
	UserId int `gorm:"column:user_id;type:INT8 NOT NULL;" json:"user_id"`
}

const (
	Small = "small"
	Large = "large"
)

type ProductRepository interface {
	// RegisterProduct 상품 등록
	RegisterProduct(ctx context.Context, product *Product) error
}

type ProductUsecase interface {
	// RegisterProduct 상품 등록
	RegisterProduct(ctx context.Context, category string, price, cost float32, name, description, barcode string, expirationData time.Time, size string, userId int) (Product, error)
}
