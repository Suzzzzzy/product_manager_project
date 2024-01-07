package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	// 유저의 고유 아이디
	Id int `gorm:"primaryKey;autoIncrement;column:id;type:INT11" json:"id"`
	// 유저의 비밀번호
	Password string `gorm:"column:password;type:VARCHAR(225) NOT NULL;" json:"password"`
	// 유저의 휴대폰 번호
	PhoneNumber string         `gorm:"column:phone_number;type:VARCHAR(20) NOT NULL;" json:"phoneNumber"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMPTZ;"`
}

func (u *User) TableName() string {
	return "users"
}

type UserRepository interface {
	// CreateUser 유저 정보 생성
	CreateUser(ctx context.Context, tx *gorm.DB, user *User) error
	// GetByPhoneNum 휴대폰 번호로 유저 조회
	GetByPhoneNum(ctx context.Context, phoneNum string) (*User, error)
	// GetByUserId 유저 고유값으로 유저 정보 조회
	GetByUserId(ctx context.Context, userId int) (*User, error)
}

type UserUsecase interface {
	// SignUp 회원가입
	SignUp(ctx context.Context, password, phoneNumber string) error
	// SignIn 로그인
	SignIn(ctx context.Context, password, phoneNumber string) (accessToken string, err error)
}
