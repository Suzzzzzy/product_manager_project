package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "root"
const DB_NAME = "product-manager-project"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var DB *gorm.DB

func InitDB() *gorm.DB {
	DB = connectDB()
	return DB
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	fmt.Println("DB connection success")

	//db.Migrator().DropTable(&model.User{})
	//db.AutoMigrate(&model.User{})

	return db
}
