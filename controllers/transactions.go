package controllers

import (
	"fmt"
	"strconv"
	"testInfotecs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendMoney выполняет перевод средств
// @Summary Отправить деньги
// @Description Переводит деньги с одного кошелька на другой
// @Accept json
// @Produce json
// @Param transaction body models.Transaction true "Данные транзакции"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/send [post]
func SendMoney(c *gin.Context, db *gorm.DB) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	if transaction.Amount <= 0 {
		c.JSON(400, gin.H{"error": "invalid value"})
		return
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		var sender, receiver models.Wallet

		if err := tx.Where("address = ?", transaction.From).First(&sender).Error; err != nil {
			c.JSON(404, gin.H{"error": "Sender wallet not found"})
			return err
		}
		if err := tx.Where("address = ?", transaction.To).First(&receiver).Error; err != nil {
			c.JSON(404, gin.H{"error": "Receiver wallet not found"})
			return err
		}

		if sender.Balance < transaction.Amount {
			c.JSON(400, gin.H{"error": "Insufficient funds"})
			return fmt.Errorf("insufficient funds")
		}

		sender.Balance -= transaction.Amount
		receiver.Balance += transaction.Amount

		if err := tx.Model(&models.Wallet{}).Where("address = ?", sender.Address).Update("balance", sender.Balance).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Wallet{}).Where("address = ?", receiver.Address).Update("balance", receiver.Balance).Error; err != nil {
			return err
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return
	}

	c.JSON(200, gin.H{"message": "Transaction successful"})
}

// GetLastTransactions возвращает N последних транзакций
// @Summary Получить последние транзакции
// @Description Возвращает N последних по времени переводов средств
// @Produce json
// @Param count query int false "Количество транзакций (по умолчанию 10)"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/transactions [get]
func GetLastTransactions(c *gin.Context, db *gorm.DB) {
	countStr := c.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countStr)
	if err != nil || count <= 0 {
		c.JSON(400, gin.H{"error": "Invalid count parameter"})
		return
	}

	var transactions []models.Transaction
	if err := db.Order("id DESC").Limit(count).Find(&transactions).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(200, transactions)
}
