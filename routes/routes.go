package routes

import (
	"blog/controllers"
	"blog/utils"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up all the routes for the blog application
func SetupRouter() *gin.Engine {
	// Create a new Gin router with default middleware
	r := gin.Default()

	// Add custom template functions
	funcMap := template.FuncMap{
		"now": func() time.Time {
			return time.Now()
		},
		"slicestr": func(s string, start, end int) string {
			if start < 0 {
				start = 0
			}
			if end > len(s) {
				end = len(s)
			}
			if start > end {
				return ""
			}
			return s[start:end]
		},
	}

	// Set the template function map
	r.SetFuncMap(funcMap)

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")

	// Get the database connection
	db := utils.DB

	// Initialize controllers
	postController := controllers.NewPostController(db)

	// Routes for blog posts
	r.GET("/", postController.Index)              // Show welcome page
	r.GET("/all-posts", postController.AllPosts)  // Show all posts
	r.GET("/posts/:id", postController.Show)      // Show a single post
	r.GET("/create", postController.CreateForm)   // Show create post form
	r.POST("/create", postController.Create)      // Create a new post
	r.GET("/edit/:id", postController.EditForm)    // Show edit post form
	r.POST("/edit/:id", postController.Edit)       // Update a post
	r.POST("/delete/:id", postController.Delete)   // Delete a post

	// Routes for comments
	r.POST("/posts/:id/comments", postController.AddComment) // Add a comment to a post

	// Test route to check if templates are working
	r.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"Title": "Test Page",
		})
	})
	
	// Test route for index template
	r.GET("/test-index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"posts": []map[string]interface{}{},
		})
	})
	
	// Test route with simple template
	r.GET("/test-simple", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test-simple.html", gin.H{
			"posts": []map[string]interface{}{},
		})
	})

	return r
}