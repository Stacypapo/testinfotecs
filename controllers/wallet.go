package controllers

import (
	"testInfotecs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetBalance возвращает баланс кошелька
// @Summary Получить баланс кошелька
// @Description Возвращает баланс по адресу кошелька
// @Produce json
// @Param address path string true "Адрес кошелька"
// @Success 200 {object} models.Wallet
// @Failure 404 {object} models.ErrorResponse
// @Router /api/wallet/{address}/balance [get]
func GetBalance(c *gin.Context, db *gorm.DB) {
	address := c.Param("address")
	var wallet models.Wallet
	if err := db.Where("address = ?", address).First(&wallet).Error; err != nil {
		c.JSON(404, gin.H{"error": "Wallet not found"})
		return
	}
	c.JSON(200, wallet)
}
