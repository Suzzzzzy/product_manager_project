package repository

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func NewMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	return gormDB, mock
}

//func TestCreateUser_Success(t *testing.T) {
//	db, mock := NewMockDB(t)
//	repo := NewUserRepository(db)
//
//	mock.ExpectBegin()
//	mock.ExpectQuery("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`password`,`phone_number`) VALUES ('2024-01-04 17:50:05.32','2024-01-04 17:50:05.32',NULL,'','010-0000-0000')") //.WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//	user := model.User{PhoneNumber: "010-0000-0000"}
//	if err := repo.CreateUser(context.Background(), user); err != nil {
//		t.Fatalf("Failed to insert user: %v", err)
//	}
//}
