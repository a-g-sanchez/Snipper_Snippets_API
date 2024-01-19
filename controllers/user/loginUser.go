package user

import (
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUser(c *gin.Context) {

	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	key := config.LoadJwtKey()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "snippr",
		"sub": user.Username,
		"id":  user.Id,
	})
	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"token":  tokenString,
		"status": "login successful",
	})
}
