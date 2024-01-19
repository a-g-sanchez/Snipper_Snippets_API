package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context, slice []User) []User {

	var newUser *User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	slice = append(slice, *newUser)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": newUser,
		"status":  "Created",
	})

	return slice
}