package usecase

import (
	"context"
	"example.com/m/src/domain"
	"example.com/m/src/domain/model"
	"example.com/m/src/domain/model/mocks"
	"example.com/m/src/utils"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	userUsecase         *userUsecase
	mockUserRepo        *mocks.UserRepository
	mockTransactionRepo *mocks.TransactionRepository
}

func (ts *UserUsecaseTestSuite) SetupTest() {
	ts.mockUserRepo = new(mocks.UserRepository)
	ts.mockTransactionRepo = new(mocks.TransactionRepository)
	ts.userUsecase = &userUsecase{
		transactionRepo: ts.mockTransactionRepo,
		userRepo:        ts.mockUserRepo,
	}
}

func TestUserUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (ts *UserUsecaseTestSuite) Test_SignUp() {
	ts.Run("회원가입 성공", func() {
		ts.mockUserRepo.On("GetByPhoneNum", mock.Anything, mock.Anything).Return(nil, nil)
		ts.mockUserRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(nil)

		result := ts.userUsecase.SignUp(context.Background(), "1234", "010-2205-0887")
		ts.NoError(result)
	})
	ts.Run("회원가입 실패 - 이미 가입한 회원 정보", func() {
		ts.SetupTest()
		ts.mockUserRepo.On("GetByPhoneNum", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockUserRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(nil)

		result := ts.userUsecase.SignUp(context.Background(), "1234", "010-0000-0000")
		ts.Equal(domain.ErrUserConflict, result)
	})
}

func (ts *UserUsecaseTestSuite) Test_SignIn() {
	password := "1234"
	hashedPassword, _ := utils.HashPassword(password)
	ts.Run("로그인 성공", func() {
		ts.mockUserRepo.On("GetByPhoneNum", mock.Anything, mock.Anything).Return(&model.User{Id: 1, Password: hashedPassword}, nil)

		result, err := ts.userUsecase.SignIn(context.Background(), password, "010-0000-0000")
		ts.NoError(err)
		ts.NotNil(result)
	})
	ts.Run("로그인 실패 - 잘못된 비밀번호", func() {
		ts.mockUserRepo.On("GetByPhoneNum", mock.Anything, mock.Anything).Return(&model.User{Id: 1, Password: hashedPassword}, nil)

		_, err := ts.userUsecase.SignIn(context.Background(), "12345", "010-0000-0000")
		ts.Equal(domain.ErrWrongPassword, err)
	})
}
