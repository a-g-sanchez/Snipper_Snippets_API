package main

import (
	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/snippet"
	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/user"
	"github.com/a-g-sanchez/Snipper_Snippets_API/middleware"
	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

func main() {

	slice := util.ParseYaml()

	var usersSlice []user.User

	router := gin.Default()

	// Snippet route group
	snippetRoutes := router.Group("/snippets")

	// // POST Route to create a new snippet and encryption middleware
	snippetRoutes.POST("/", middleware.SnippetEncryption(), func(c *gin.Context) {
		slice = snippet.AddNewSnippet(c, slice)
	})

	// GET all snippets or a snippet based on a query for a certain language
	snippetRoutes.GET("/", func(c *gin.Context) {
		snippet.GetAllSnippets(c, slice)

	})

	// GET a snippet by ID
	snippetRoutes.GET("/:id", func(c *gin.Context) {
		snippet.GetSnippetById(c, slice)
	})

	// User route group
	userRoutes := router.Group("/users")

	// POST / register a new user with a hashed and salted password
	userRoutes.POST("/register", middleware.HashPassword(), func(c *gin.Context) {
		usersSlice = user.RegisterUser(c, usersSlice)
	})

	// POST / login a user
	// Need to look into req.user
	userRoutes.POST("/login", middleware.CompareHash(&usersSlice), func(c *gin.Context) {
		user.LoginUser(c)
	})

	router.Run()
}
