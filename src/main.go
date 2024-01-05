package main

import (
	"example.com/m/src/adapter/http"
	"example.com/m/src/repository"
	"example.com/m/src/repository/mysql"
	"example.com/m/src/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ht "net/http"
)

func main() {
	db := mysql.InitDB()
	if db == nil {
		panic(db)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(ht.StatusOK, "Hello, World!")
	})

	transactionRepo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(transactionRepo, userRepo)
	http.NewUserHandler(router, userUsecase)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
