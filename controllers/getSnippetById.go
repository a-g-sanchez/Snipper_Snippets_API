package controllers

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

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": foundSnippet,
	})

}
