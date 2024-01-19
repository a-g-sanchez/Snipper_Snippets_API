package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-g-sanchez/Snipper_Snippets_API/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {

		key := config.LoadJwtKey()
		// WILL ADD ERROR HANDLING IF NO HEADER OR
		// IF AUTH TYPE IS NOT BEARER
		header := c.Request.Header["Authorization"]

		headerString := header[0]

		token := strings.Split(headerString, " ")[1]

		_, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized!",
			})
			c.Abort()
			return
		}

		// claims, ok := validToken.Claims.(jwt.MapClaims)
		// if !ok {
		// 	c.IndentedJSON(http.StatusUnauthorized, gin.H{
		// 		"error": "unauthorized",
		// 	})
		// 	c.Abort()
		// }
		// if claims["exp"] != nil {
		// 	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		// 		c.IndentedJSON(http.StatusUnauthorized, gin.H{
		// 			"error": "Token expired",
		// 		})
		// 	}
		// }

	}
}
