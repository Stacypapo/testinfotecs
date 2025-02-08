package main

import (
	"fmt"
	"math/rand"
	_ "testInfotecs/docs"
	"testInfotecs/models"
	"testInfotecs/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const initialBalance = 100.0
const walletCount = 10

// generateRandomAddress создает случайный адрес для кошелька
func generateRandomAddress() string {
	letters := "abcdef0123456789"
	address := make([]byte, 64)
	for i := range address {
		address[i] = letters[rand.Intn(len(letters))]
	}
	return string(address)
}

// seedWallets проверяет наличие кошельков в БД
// и в случае если они отсутствуют создает {walletCount} кошельков
// с {initialBalance} на счету
func seedWallets(db *gorm.DB) {
	var count int64
	db.Model(&models.Wallet{}).Count(&count)
	if count == 0 {
		for i := 0; i < walletCount; i++ {
			wallet := models.Wallet{
				Address: generateRandomAddress(),
				Balance: initialBalance,
			}
			db.Create(&wallet)
		}
		fmt.Println("10 кошельков созданы")
	} else {
		fmt.Println("Кошельки уже существуют")
	}
}

// @title Payment System API
// @version 1.0
// @description API для управления транзакциями и кошельками
// @host localhost:8080
// @BasePath /
func main() {
	dsn := "user=postgres password=12345 dbname=test host=10.8.1.2 port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connected!")

	models.AutoMigrateWallets(db)
	models.AutoMigrateTransactions(db)

	seedWallets(db)

	r := gin.Default()
	routes.SetupRoutes(r, db)

	r.Run("0.0.0.0:8080")
}
