package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID     int     `json:"id"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func AutoMigrateTransactions(db *gorm.DB) {
	db.AutoMigrate(&Transaction{})
}
