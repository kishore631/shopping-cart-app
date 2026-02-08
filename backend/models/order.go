package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Order struct {
	ID     uint
	UserID uint
	Items  []Item
}

var Orders []Order

func CreateOrder(c *gin.Context) {
	userID := uint(1)
	var orderItems []Item

	for _, c := range Cart {
		if c.UserID == userID {
			orderItems = append(orderItems, c.Item)
		}
	}

	if len(orderItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}

	order := Order{
		ID:     uint(len(Orders) + 1),
		UserID: userID,
		Items:  orderItems,
	}

	Orders = append(Orders, order)
	Cart = []CartItem{} // clear cart

	c.JSON(http.StatusOK, gin.H{"message": "Order placed", "order": order})
}

func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, Orders)
}
