package controllers

import (
	"net/http"
	"strings"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func GetAllSnippets(c *gin.Context, slice []util.Snippet) {
	lang := c.Query("lang")

	if lang != "" {
		var filteredSnippets []util.Snippet

		for _, snippet := range slice {
			if strings.EqualFold(lang, snippet.Language) {
				filteredSnippets = append(filteredSnippets, snippet)
			}
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"message": filteredSnippets,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": slice,
	})
}
