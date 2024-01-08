package usecase

import (
	"context"
	"example.com/m/src/domain"
	"example.com/m/src/domain/model"
	"example.com/m/src/domain/model/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type ProductUsecaseTestSuite struct {
	suite.Suite
	productUsecase      *productUsecase
	mockTransactionRepo *mocks.TransactionRepository
	mockUserRepo        *mocks.UserRepository
	mockProductRepo     *mocks.ProductRepository
}

func (ts *ProductUsecaseTestSuite) SetupTest() {
	ts.mockTransactionRepo = new(mocks.TransactionRepository)
	ts.mockUserRepo = new(mocks.UserRepository)
	ts.mockProductRepo = new(mocks.ProductRepository)
	ts.productUsecase = &productUsecase{
		transactionRepo: ts.mockTransactionRepo,
		userRepo:        ts.mockUserRepo,
		productRepo:     ts.mockProductRepo,
	}
}

func TestProductUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(ProductUsecaseTestSuite))
}

func (ts *ProductUsecaseTestSuite) Test_RegisterProduct() {
	product := &model.Product{
		Id:     1,
		Name:   "붕어빵",
		UserId: 1,
	}
	ts.Run("상품 등록 성공", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("RegisterProduct", mock.Anything, mock.Anything, mock.Anything).Return(product, nil)
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(nil)

		_, err := ts.productUsecase.RegisterProduct(context.Background(), &model.Product{Name: "붕어빵"}, 1)
		ts.NoError(err)
	})
}

func (ts *ProductUsecaseTestSuite) Test_GetByProductId() {
	product := &model.Product{
		Id:     1,
		Name:   "붕어빵",
		UserId: 1,
	}
	ts.Run("상품 조회 성공", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("GetByProductId", mock.Anything, mock.Anything).Return(product, nil)

		result, err := ts.productUsecase.GetByProductId(context.Background(), product.Id, 1)
		ts.NoError(err)
		ts.Equal("붕어빵", result.Name)
	})
	ts.Run("상품 조회 실패 - 상품 접근 권한 없음", func() {
		ts.SetupTest()
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 2}, nil)
		ts.mockProductRepo.On("GetByProductId", mock.Anything, mock.Anything).Return(product, nil)

		_, err := ts.productUsecase.GetByProductId(context.Background(), product.Id, 2)
		ts.Equal(domain.ErrInvalidUser, err)
	})
}

func (ts *ProductUsecaseTestSuite) Test_UpdateProduct() {
	product := &model.Product{
		Id:     1,
		Name:   "붕어빵",
		UserId: 1,
	}
	updateInfo := map[string]interface{}{
		"name": "슈크림 붕어빵",
	}
	ts.Run("상품 조회 성공", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("GetByProductId", mock.Anything, mock.Anything).Return(product, nil)
		product.Name = "슈크림 붕어빵"
		ts.mockProductRepo.On("UpdateProduct", mock.Anything, mock.Anything, mock.Anything).Return(product, nil)

		result, err := ts.productUsecase.UpdateProduct(context.Background(), product.Id, 1, updateInfo)
		ts.NoError(err)
		ts.Equal(updateInfo["name"], result.Name)
	})
}

func (ts *ProductUsecaseTestSuite) Test_DeleteProduct() {
	product := &model.Product{
		Id:     1,
		Name:   "붕어빵",
		UserId: 1,
	}
	ts.Run("상품 삭제 성공", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("GetByProductId", mock.Anything, mock.Anything).Return(product, nil)
		ts.mockProductRepo.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil)

		err := ts.productUsecase.DeleteProduct(context.Background(), product.Id, 1)
		ts.NoError(err)
	})
	ts.Run("상품 삭제 실패 - 없는 상품", func() {
		ts.SetupTest()
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("GetByProductId", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound)
		ts.mockProductRepo.On("DeleteProduct", mock.Anything, mock.Anything).Return(nil)

		err := ts.productUsecase.DeleteProduct(context.Background(), product.Id, 1)
		ts.Equal(domain.ErrProductNotFound, err)
	})

}

func generateTestProducts(n int) []model.Product {
	var productList []model.Product
	for i := 1; i <= n; i++ {
		prod := model.Product{
			Id:     i,
			Name:   "붕어빵",
			UserId: 1,
		}
		productList = append(productList, prod)
	}
	return productList
}

func (ts *ProductUsecaseTestSuite) Test_GetProductList() {
	productList := generateTestProducts(10)
	totalPage := 2

	ts.Run("상품 리스트 조회", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("GetTotalProductCount", mock.Anything, mock.Anything).Return(12, nil)
		ts.mockProductRepo.On("GetProductList", mock.Anything, mock.Anything, mock.Anything).Return(productList, nil)

		result, page, err := ts.productUsecase.GetProductList(context.Background(), 1, 1)
		ts.NoError(err)
		ts.Equal(totalPage, page)
		ts.Equal(10, len(result))
	})
	ts.Run("상품 리스트 조회 - 없는 페이지 조회", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("GetTotalProductCount", mock.Anything, mock.Anything).Return(12, nil)
		ts.mockProductRepo.On("GetProductList", mock.Anything, mock.Anything, mock.Anything).Return(productList, nil)

		result, _, err := ts.productUsecase.GetProductList(context.Background(), 1, 3)
		ts.NoError(err)
		ts.Equal(0, len(result))
	})
}

func (ts *ProductUsecaseTestSuite) Test_FindProductByName() {
	productList := generateTestProducts(3)
	ts.Run("상품 검색 조회", func() {
		ts.mockUserRepo.On("GetByUserId", mock.Anything, mock.Anything).Return(&model.User{Id: 1}, nil)
		ts.mockProductRepo.On("FindProductByName", mock.Anything, mock.Anything, mock.Anything).Return(productList, nil)

		result, err := ts.productUsecase.FindProductByName(context.Background(), 1, "ㅂㅇㅃ")
		ts.NoError(err)
		ts.NotNil(result)
	})
}
