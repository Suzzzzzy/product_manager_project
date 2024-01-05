package usecase

import (
	"context"
	"errors"
	"example.com/m/src/domain"
	"example.com/m/src/domain/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type userUsecase struct {
	transactionRepo model.TransactionRepository
	userRepo        model.UserRepository
}

func NewUserUsecase(transactionRepo model.TransactionRepository, userRepo model.UserRepository) model.UserUsecase {
	return &userUsecase{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

func (u userUsecase) CreateUser(ctx context.Context, password, phoneNumber string) error {
	// 가입내역 확인
	alreadyUser, err := u.userRepo.GetByPhoneNum(ctx, phoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrInternalServerError
	}
	if alreadyUser != nil {
		return domain.ErrConflict
	}
	// 비밀번호 해쉬
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Printf("failed to GenerateFromPassword: %v", err)
		return domain.ErrInternalServerError
	}
	user := &model.User{
		Password:    string(hashedPassword),
		PhoneNumber: phoneNumber,
	}

	if txErr := u.transactionRepo.Transaction(func(tx *gorm.DB) error {
		return u.userRepo.CreateUser(ctx, tx, user)
	}); txErr != nil {
		return domain.ErrInternalServerError
	}
	return nil
}
