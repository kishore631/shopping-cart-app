package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Items []Item

type Item struct {
	ID    uint
	Name  string
	Price float64
}

func CreateItem(c *gin.Context) {
	var item Item
	c.BindJSON(&item)

	item.ID = uint(len(Items) + 1)
	Items = append(Items, item)

	c.JSON(http.StatusOK, item)
}

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, Items)
}
