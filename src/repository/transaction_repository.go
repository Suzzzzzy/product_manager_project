package repository

import (
	"example.com/m/src/domain/model"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) model.TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) Transaction(txFn func(tx *gorm.DB) error) error {
	return t.db.Transaction(txFn)
}
