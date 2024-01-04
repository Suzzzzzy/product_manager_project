package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	ht "net/http"
)

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", "root", "root", "mysql", "3306", "product_manager_project")
	_, err := sql.Open(`mysql`, connection)
	if err != nil {
		fmt.Println("DB connection failed:", err)
		return
	}
	fmt.Println("DB connection success")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(ht.StatusOK, "Hello, World!")
	})

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
