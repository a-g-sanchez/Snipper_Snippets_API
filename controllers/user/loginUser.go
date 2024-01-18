package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {

	// fmt.Println(c.Request.Body)
	// var user *User
	// for _, users := range slice {
	// 	if users.Username == username {
	// 		user = &users
	// 		break
	// 	}
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{
	// 		"message": fmt.Sprintf("No user with the name %s", username),
	// 	})
	// }

	// err := util.ComparePasswords(user.Password, password)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{
	// 		"message": "Passwords do not match try again",
	// 	})
	// } else {
	// 	c.IndentedJSON(http.StatusOK, gin.H{
	// 		"user": user.Username,
	// 	})
	// }

	c.IndentedJSON(http.StatusCreated, gin.H{
		"tokin":  "token goes here",
		"status": "login successful",
	})
}
