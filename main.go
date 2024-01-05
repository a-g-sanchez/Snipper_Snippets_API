package main

import (
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func main() {

	slice := util.ParseYaml()

	router := gin.Default()

	// ROUTES
	router.GET("/snippets", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": slice,
		})
	})

	router.POST("/snippets", func(c *gin.Context) {
		var newSnippet *util.Snippet
		if err := c.ShouldBindJSON(&newSnippet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		slice = append(slice, *newSnippet)

		c.IndentedJSON(http.StatusCreated, gin.H{
			"message": slice,
		})
	})

	router.Run()
}
