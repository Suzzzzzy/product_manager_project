package usecase

import (
	"context"
	"errors"
	"example.com/m/src/domain"
	"example.com/m/src/domain/model"
	"gorm.io/gorm"
)

type productUsecase struct {
	transactionRepo model.TransactionRepository
	userRepo        model.UserRepository
	productRepo     model.ProductRepository
}

func NewProductUsecase(transactionRepo model.TransactionRepository, userRepo model.UserRepository, productRepo model.ProductRepository) model.ProductUsecase {
	return &productUsecase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		productRepo:     productRepo,
	}
}

// checkUserValid 유저의 유효성을 검증
func (p *productUsecase) checkUserValid(ctx context.Context, userId int) (*model.User, error) {
	user, err := p.userRepo.GetByUserId(ctx, userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrInternalServerError
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

// checkProductValid 상품의 유효성을 검증
func (p *productUsecase) checkProductValid(ctx context.Context, productId, userId int) (*model.Product, error) {
	product, err := p.productRepo.GetByProductId(ctx, productId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrInternalServerError
	}
	if product == nil {
		return nil, domain.ErrProductNotFound
	}
	if product.UserId != userId {
		return nil, domain.ErrInvalidUser
	}
	return product, nil
}

func (p *productUsecase) RegisterProduct(ctx context.Context, product *model.Product, userId int) (*model.Product, error) {
	// 유저 검증
	_, err := p.checkUserValid(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 상품 정보 저장
	var newProduct *model.Product
	if txErr := p.transactionRepo.Transaction(func(tx *gorm.DB) error {
		createdProduct, err := p.productRepo.RegisterProduct(ctx, tx, product)
		if err != nil {
			return err
		}
		newProduct = createdProduct
		return nil
	}); txErr != nil {
		return nil, domain.ErrInternalServerError
	}

	return newProduct, nil
}

func (p *productUsecase) GetByProductId(ctx context.Context, productId, userId int) (*model.Product, error) {
	// 유저 검증
	_, err := p.checkUserValid(ctx, userId)
	if err != nil {
		return nil, err
	}
	// 상품 검증
	product, err := p.checkProductValid(ctx, productId, userId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *productUsecase) UpdateProduct(ctx context.Context, productId, userId int, updateInfo map[string]interface{}) (*model.Product, error) {
	// 유저 검증
	_, err := p.checkUserValid(ctx, userId)
	if err != nil {
		return nil, err
	}
	// 상품 검증
	product, err := p.checkProductValid(ctx, productId, userId)
	if err != nil {
		return nil, err
	}
	// 상품 정보 수정(특정 칼럼만)
	updatedProduct, err := p.productRepo.UpdateProduct(ctx, product, updateInfo)
	if err != nil {
		return nil, err
	}
	return updatedProduct, nil
}

func (p *productUsecase) DeleteProduct(ctx context.Context, productId, userId int) error {
	// 유저 검증
	_, err := p.checkUserValid(ctx, userId)
	if err != nil {
		return err
	}
	// 상품 검증
	product, err := p.checkProductValid(ctx, productId, userId)
	if err != nil {
		return err
	}
	// 상품 삭제
	err = p.productRepo.DeleteProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
