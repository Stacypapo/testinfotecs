package routes

import (
	"testInfotecs/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	//r.HandleFunc("/api/send", handlers.SendMoney).Methods("POST")
	//r.HandleFunc("/api/transactions", handlers.GetLastTransactions).Methods("GET")
	//r.HandleFunc("/api/wallet/{address}/balance", handlers.GetBalance).Methods("GET")
	//r.LoadHTMLGlob("templates/*")
	//r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "abc")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/api/send", func(c *gin.Context) {
		controllers.SendMoney(c, db)
	})

	r.GET("/api/transactions", func(c *gin.Context) {
		controllers.GetLastTransactions(c, db)
	})

	r.GET("/api/wallet/:address/balance", func(c *gin.Context) {
		controllers.GetBalance(c, db)
	})

}
