package repository

import (
	"context"
	"example.com/m/src/domain/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, tx *gorm.DB, user *model.User) error {
	return tx.WithContext(ctx).Model(&model.User{}).Create(user).Error
}

func (u *userRepository) GetByPhoneNum(ctx context.Context, phoneNumber string) (*model.User, error) {
	var user *model.User
	err := u.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
