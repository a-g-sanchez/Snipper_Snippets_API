package main

import (
	"fmt"
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers"
	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {

	slice := util.ParseYaml()

	var usersSlice []User

	router := gin.Default()

	// SNIPPET ROUTES
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

	//USER ROUTES
	// POST / create a new user with a hashed and salted password
	router.POST("/users/", func(c *gin.Context) {

		var newUser *User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		hashedPassword, err := util.HashPassword(newUser.Password)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		newUser.Password = hashedPassword

		usersSlice = append(usersSlice, *newUser)

		c.IndentedJSON(http.StatusCreated, gin.H{
			"message": newUser,
			"status":  "Created",
		})
	})

	// GET a user
	router.GET("/users", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		var user *User
		for _, users := range usersSlice {
			if users.Username == username {
				user = &users
				break
			}
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("No user with the name %s", username),
			})
		}

		err := util.ComparePasswords(user.Password, password)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "Passwords do not match try again",
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"user": user.Username,
			})
		}

		// fmt.Println(user)

	})

	router.Run()
}
