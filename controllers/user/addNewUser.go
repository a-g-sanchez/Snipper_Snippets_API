package user

import (
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AddNewUser(c *gin.Context, slice []User) []User {

	var newUser *User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	newUser.Password = hashedPassword

	slice = append(slice, *newUser)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": newUser,
		"status":  "Created",
	})

	return slice
}
