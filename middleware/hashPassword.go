package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/user"
	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func HashPassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *user.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		user.Password = hashedPassword

		serializedData, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(serializedData))
		c.Request.ContentLength = int64(len(serializedData))

		c.Next()

	}
}
