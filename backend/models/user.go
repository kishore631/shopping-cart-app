package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint
	Username string
	Password string
	Token    string
}

var Users []User

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashed)
	user.ID = uint(len(Users) + 1)

	Users = append(Users, user)
	c.JSON(http.StatusOK, user)
}

func LoginUser(c *gin.Context) {
	var input User
	c.BindJSON(&input)

	for i, u := range Users {
		if u.Username == input.Username {
			err := bcrypt.CompareHashAndPassword(
				[]byte(u.Password),
				[]byte(input.Password),
			)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				return
			}

			Users[i].Token = "sample-token"
			c.JSON(http.StatusOK, gin.H{"token": Users[i].Token})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
}
