package middleware

import (
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/user"
	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func CompareHash(userSlice *[]user.User) gin.HandlerFunc {
	return func(c *gin.Context) {

		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")

		var foundUser *user.User

		// *** TEMP LOGIC TO FIND A USER FROM THE SLICE UNTIL DB IS ADDED ***
		for _, user := range *userSlice {
			if username == user.Username {
				foundUser = &user
				break
			}
		}

		if foundUser == nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "User not found!",
			})
			c.Abort()

		}

		// fmt.Println(foundUser)
		// // *** SHOULD BE END OF TEMP CODE ***
		// // The rest shouldnt have to change

		err := util.ComparePasswords(foundUser.Password, password)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid password",
			})
			c.Abort()
		}

		c.Next()
	}

}
