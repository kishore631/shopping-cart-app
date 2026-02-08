package main

import (
	"shopping-cart/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// ✅ CORS FIX
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.POST("/users", models.CreateUser)
	r.POST("/users/login", models.LoginUser)

	r.POST("/items", models.CreateItem)
	r.GET("/items", models.GetItems)

	r.POST("/carts", models.AddToCart)
	r.GET("/carts", models.GetCart)

	r.POST("/orders", models.CreateOrder)
	r.GET("/orders", models.GetOrders)

	r.Run(":8080")
}
