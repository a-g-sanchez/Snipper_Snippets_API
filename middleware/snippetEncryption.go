package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func SnippetEncryption() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newSnippet *util.Snippet

		if err := c.ShouldBindJSON(&newSnippet); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		encryptedData, err := util.Encrypt([]byte(newSnippet.Code))
		if err != nil {
			panic(err)
		}

		newSnippet.Code = encryptedData

		serializedData, err := json.Marshal(newSnippet)
		if err != nil {
			panic(err)
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(serializedData))
		c.Request.ContentLength = int64(len(serializedData))
		c.Next()
	}

}
