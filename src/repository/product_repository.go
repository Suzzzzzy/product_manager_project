package repository

import (
	"context"
	"example.com/m/src/domain/model"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) model.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (p *productRepository) RegisterProduct(ctx context.Context, tx *gorm.DB, product *model.Product) (*model.Product, error) {
	if err := tx.WithContext(ctx).Model(&model.Product{}).Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepository) GetByProductId(ctx context.Context, productId int) (*model.Product, error) {
	var product *model.Product
	err := p.db.WithContext(ctx).Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
