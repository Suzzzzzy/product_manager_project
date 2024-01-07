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
	var product model.Product
	err := p.db.WithContext(ctx).Where("id = ?", productId).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) UpdateProduct(ctx context.Context, product *model.Product, updateInfo map[string]interface{}) (*model.Product, error) {
	if err := p.db.WithContext(ctx).Model(product).Updates(updateInfo).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productRepository) DeleteProduct(ctx context.Context, product *model.Product) error {
	if err := p.db.WithContext(ctx).Delete(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *productRepository) GetProductList(ctx context.Context, userId, page int) ([]model.Product, error) {
	var productList []model.Product
	pageSize := 10
	offset := (page - 1) * pageSize
	err := p.db.WithContext(ctx).Where("user_id = ?", userId).Order("created_at desc").Offset(offset).Limit(pageSize).Find(&productList).Error
	return productList, err
}

func (p *productRepository) GetTotalProductCount(ctx context.Context, userId int) (int, error) {
	var count int64
	err := p.db.WithContext(ctx).Model(&model.Product{}).Where("user_id = ?", userId).Count(&count).Error
	return int(count), err
}
