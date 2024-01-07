package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	gorm.Model
	// 상품의 고유 아이디
	Id int `gorm:"primaryKey;autoIncrement;column:id;type:INT11" json:"id"`
	// 카테고리
	Category string `gorm:"column:category;type:VARCHAR NOT NULL;" json:"category"`
	// 가격
	Price float32 `gorm:"column:price;type:FLOAT NOT NULL;" json:"price"`
	// 원가
	Cost float32 `gorm:"column:cost;type:FLOAT NOT NULL;" json:"cost"`
	// 이름
	Name string `gorm:"column:name;type:VARCHAR NOT NULL;" json:"name"`
	// 설명
	Description string `gorm:"column:description;type:VARCHAR NOT NULL;" json:"description"`
	// 바코드
	Barcode string `gorm:"column:barcode;type:VARCHAR NOT NULL;" json:"barcode"`
	// 유통기한
	ExpirationDate time.Time `gorm:"column:expiration_date;type:DATETIME NOT NULL;" json:"expiration_date"`
	// 사이즈
	Size string `gorm:"column:size;type:VARCHAR NOT NULL;" json:"size"`
	// 상품을 등록한 유저의 고유값
	UserId int `gorm:"column:user_id;type:INT8 NOT NULL;" json:"user_id"`

	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMPTZ;"`
}

func (p *Product) TableName() string {
	return "product"
}

const (
	Small = "small"
	Large = "large"
)

type ProductRepository interface {
	// RegisterProduct 상품 등록
	RegisterProduct(ctx context.Context, tx *gorm.DB, product *Product) (*Product, error)
	// GetByProductId 상품 조회
	GetByProductId(ctx context.Context, productId int) (*Product, error)
	// UpdateProduct 상품 정보 수정
	UpdateProduct(ctx context.Context, product *Product, updateInfo map[string]interface{}) (*Product, error)
	// DeleteProduct 상품 단일 삭제
	DeleteProduct(ctx context.Context, product *Product) error
}

type ProductUsecase interface {
	// RegisterProduct 상품 등록
	RegisterProduct(ctx context.Context, product *Product, userId int) (*Product, error)
	// GetByProductId 상품 단일 조회
	GetByProductId(ctx context.Context, productId, userId int) (*Product, error)
	// UpdateProduct 상품 정보 수정
	UpdateProduct(ctx context.Context, productId, userId int, updateInfo map[string]interface{}) (*Product, error)
	// DeleteProduct 상품 단일 삭제
	DeleteProduct(ctx context.Context, productId, userId int) error
}
