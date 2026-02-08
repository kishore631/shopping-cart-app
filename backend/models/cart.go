package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartItem struct {
	UserID uint
	Item   Item
}

var Cart []CartItem

// 🔹 Extract user from token (TEMP SIMPLE VERSION)
func getUserIDFromToken(c *gin.Context) uint {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		return 0
	}
	// For now: single user = ID 1
	return 1
}

func AddToCart(c *gin.Context) {
	var body struct {
		ItemID uint `json:"item_id"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := getUserIDFromToken(c)

	for _, item := range Items {
		if item.ID == body.ItemID {
			Cart = append(Cart, CartItem{
				UserID: userID,
				Item:   item,
			})
			c.JSON(http.StatusOK, gin.H{"message": "Added to cart"})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
}

func GetCart(c *gin.Context) {
	userID := getUserIDFromToken(c)

	var userCart []CartItem
	for _, c := range Cart {
		if c.UserID == userID {
			userCart = append(userCart, c)
		}
	}

	c.JSON(http.StatusOK, userCart)
}
