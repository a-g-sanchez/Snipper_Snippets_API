package main

import (
	"net/http"
	"strconv"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func main() {

	slice := util.ParseYaml()

	router := gin.Default()

	// ROUTES
	// GET all snippets
	router.GET("/snippets", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": slice,
		})
	})

	// GET a snippet by ID
	router.GET("/snippets/:id", func(c *gin.Context) {
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
	})

	// POST / create a new snippet
	router.POST("/snippets", func(c *gin.Context) {
		var newSnippet *util.Snippet
		if err := c.ShouldBindJSON(&newSnippet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		slice = append(slice, *newSnippet)

		c.IndentedJSON(http.StatusCreated, gin.H{
			"message": newSnippet,
		})
	})

	router.Run()
}
