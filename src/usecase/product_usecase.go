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

func (p productUsecase) RegisterProduct(ctx context.Context, product *model.Product, userId int) (*model.Product, error) {
	// 유저 검증
	user, err := p.userRepo.GetByUserId(ctx, userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, domain.ErrInternalServerError
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
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
