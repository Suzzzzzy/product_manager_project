package model

import (
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Transaction(txFn func(tx *gorm.DB) error) error
}
