package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	Address string  `json:"address"`
	Balance float64 `json:"balance"`
}

func AutoMigrateWallets(db *gorm.DB) {
	db.AutoMigrate(&Wallet{})
}
