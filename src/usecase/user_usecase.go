package usecase

import (
	"context"
	"errors"
	"example.com/m/src/domain"
	"example.com/m/src/domain/model"
	"example.com/m/src/utils"
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

func (u userUsecase) SignUp(ctx context.Context, password, phoneNumber string) error {
	// 가입내역 확인
	alreadyUser, err := u.userRepo.GetByPhoneNum(ctx, phoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.ErrInternalServerError
	}
	if alreadyUser != nil {
		return domain.ErrUserConflict
	}
	// 비밀번호 해쉬
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Printf("failed to GenerateFromPassword: %v", err)
		return domain.ErrInternalServerError
	}
	user := &model.User{
		Password:    hashedPassword,
		PhoneNumber: phoneNumber,
	}

	if txErr := u.transactionRepo.Transaction(func(tx *gorm.DB) error {
		return u.userRepo.CreateUser(ctx, tx, user)
	}); txErr != nil {
		return domain.ErrInternalServerError
	}
	return nil
}

func (u userUsecase) SignIn(ctx context.Context, inputPassword, phoneNumber string) (string, error) {
	// 가입내역 확인
	user, err := u.userRepo.GetByPhoneNum(ctx, phoneNumber)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", domain.ErrInternalServerError
	}
	if user == nil {
		return "", domain.ErrUserNotFound
	}
	// 비밀번호 검증
	res := utils.CheckPasswordHash(user.Password, inputPassword)
	if !res {
		return "", domain.ErrWrongPassword
	}
	// 토큰 발행
	accessToken, err := utils.CreateJWT(user.PhoneNumber, int(user.ID))
	if err != nil {
		log.Printf("failed to Generate AccessToken")
		return "", domain.ErrInternalServerError
	}

	return accessToken, nil
}
