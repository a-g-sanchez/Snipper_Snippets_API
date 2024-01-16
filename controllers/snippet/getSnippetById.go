package snippet

import (
	"net/http"
	"strconv"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func GetSnippetById(c *gin.Context, slice []util.Snippet) {
	id := c.Param("id")
	snippetId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	var foundSnippet *util.Snippet

	for _, snippet := range slice {
		if snippet.Id == snippetId {
			foundSnippet = &snippet
			break
		}
	}

	// Only works when a snippet has encrypted data
	// Waiting for db implementation and all snippets
	// Will have encrypted data
	decrypted, err := util.Decrypt(foundSnippet.Code)
	if err != nil {
		panic(err)
	}

	foundSnippet.Code = string(decrypted)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": foundSnippet,
	})

}
