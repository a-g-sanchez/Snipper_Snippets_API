package user

import (
	"fmt"
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, slice []User) {
	username := c.Query("username")
	password := c.Query("password")

	var user *User
	for _, users := range slice {
		if users.Username == username {
			user = &users
			break
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("No user with the name %s", username),
		})
	}

	err := util.ComparePasswords(user.Password, password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Passwords do not match try again",
		})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"user": user.Username,
		})
	}

}
