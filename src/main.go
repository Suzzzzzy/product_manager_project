package main

import (
	"context"
	"errors"
	"example.com/m/src/adapter/http"
	"example.com/m/src/repository"
	"example.com/m/src/repository/mysql"
	"example.com/m/src/usecase"
	"github.com/gin-gonic/gin"
	ht "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(transactionRepo, userRepo, productRepo)
	http.NewProductHandler(router, productUsecase)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	server := &ht.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, ht.ErrServerClosed) {
			panic(err)
		}
	}()

	<-stop
	println("\nShutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	//todo db close 추가

	println("Shutdown complete")
}
