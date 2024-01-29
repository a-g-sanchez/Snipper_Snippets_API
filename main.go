package main

import (
	"encoding/gob"
	"log"
	"net/http"

	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/app/callback"
	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/app/login"
	"github.com/a-g-sanchez/Snipper_Snippets_API/controllers/snippet"
	"github.com/a-g-sanchez/Snipper_Snippets_API/middleware"
	"github.com/a-g-sanchez/Snipper_Snippets_API/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func New(auth *util.Authenticator, snippetSlice []util.Snippet) *gin.Engine {

	router := gin.Default()

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	// Direct user to login
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))

	// Snippet route group
	snippetRoutes := router.Group("/snippets")

	// // POST Route to create a new snippet and encryption middleware
	snippetRoutes.POST("/", middleware.SnippetEncryption(), func(c *gin.Context) {
		snippetSlice = snippet.AddNewSnippet(c, snippetSlice)
	})

	// Auth middleware added to verify a user is logged in
	// GET all snippets or a snippet based on a query for a certain language
	snippetRoutes.GET("/", middleware.IsAuthenticated, func(c *gin.Context) {
		snippet.GetAllSnippets(c, snippetSlice)

	})

	// GET a snippet by ID
	snippetRoutes.GET("/:id", func(c *gin.Context) {
		snippet.GetSnippetById(c, snippetSlice)
	})

	return router
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
}

func main() {

	slice := util.ParseYaml()

	auth, err := util.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := New(auth, slice)

	log.Print("Server listening on http://localhost:8080")
	if err := http.ListenAndServe("0.0.0.0:8080", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

}
