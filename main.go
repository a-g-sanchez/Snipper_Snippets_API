package main

import (
	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers"
	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func main() {

	slice := util.ParseYaml()

	router := gin.Default()

	// ROUTES
	// GET all snippets or a snippet based on a query for a certain language
	router.GET("/snippets", func(c *gin.Context) {
		controllers.GetAllSnippets(c, slice)
	})

	// GET a snippet by ID
	router.GET("/snippets/:id", func(c *gin.Context) {
		controllers.GetSnippetById(c, slice)
	})

	// POST / create a new snippet
	router.POST("/snippets", func(c *gin.Context) {
		slice = controllers.AddNewSnippet(c, slice)
	})

	router.Run()
}
