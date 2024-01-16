package snippet

import (
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func AddNewSnippet(c *gin.Context, slice []util.Snippet) []util.Snippet {

	var newSnippet *util.Snippet

	if err := c.ShouldBindJSON(&newSnippet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	encrypted, err := util.Encrypt([]byte(newSnippet.Code))
	if err != nil {
		panic(err)
	}

	newSnippet.Code = encrypted

	slice = append(slice, *newSnippet)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": newSnippet,
		"status":  "Created",
	})

	return slice
}
