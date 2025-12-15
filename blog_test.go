package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"blog/models"
	"blog/routes"
	"blog/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestServer sets up a test server for testing
func setupTestServer() *gin.Engine {
	// Initialize database
	utils.InitDB()
	utils.AutoMigrateDB(&models.Post{}, &models.Comment{})

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Setup router
	r := routes.SetupRouter()

	return r
}

// TestBlogRoutes tests all blog routes
func TestBlogRoutes(t *testing.T) {
	// Setup test server
	r := setupTestServer()

	// Test cases
	testCases := []struct {
		name       string
		method     string
		path       string
		body       string
		statusCode int
	}{{
		name:       "Test Index Route",
		method:     "GET",
		path:       "/",
		statusCode: http.StatusOK,
	}, {
		name:       "Test Create Form Route",
		method:     "GET",
		path:       "/create",
		statusCode: http.StatusOK,
	}}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var req *http.Request
			var err error

			if tc.method == "POST" {
				req, err = http.NewRequest(tc.method, tc.path, bytes.NewBufferString(tc.body))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			} else {
				req, err = http.NewRequest(tc.method, tc.path, nil)
			}

			assert.NoError(t, err)

			// Record the response
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			// Check status code
			assert.Equal(t, tc.statusCode, w.Code)
			fmt.Printf("%s: %d OK\n", tc.name, w.Code)
		})
	}
}

// TestBlogCRUD tests CRUD operations for blog posts
func TestBlogCRUD(t *testing.T) {
	// Setup test server
	r := setupTestServer()

	// Test 1: Create a new post
	fmt.Println("\n=== Testing Create Post ===")
	createBody := "title=Test Post&author=Test Author&content=This is a test post content."
	createReq, _ := http.NewRequest("POST", "/create", bytes.NewBufferString(createBody))
	createReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	createW := httptest.NewRecorder()
	r.ServeHTTP(createW, createReq)

	assert.Equal(t, http.StatusFound, createW.Code)
	fmt.Printf("Create Post: %d OK\n", createW.Code)

	// Get the post ID from the database
	var postID uint
	db := utils.DB
	var post models.Post
	err := db.Last(&post).Error
	assert.NoError(t, err)
	postID = post.ID
	fmt.Printf("Created Post ID: %d\n", postID)

	// Test 2: Get the post detail
	fmt.Println("\n=== Testing Get Post Detail ===")
	getReq, _ := http.NewRequest("GET", fmt.Sprintf("/posts/%d", postID), nil)
	getW := httptest.NewRecorder()
	r.ServeHTTP(getW, getReq)

	assert.Equal(t, http.StatusOK, getW.Code)
	fmt.Printf("Get Post Detail: %d OK\n", getW.Code)

	// Test 3: Edit the post
	fmt.Println("\n=== Testing Edit Post ===")
	editBody := "title=Updated Test Post&author=Updated Author&content=This is updated test post content."
	editReq, _ := http.NewRequest("POST", fmt.Sprintf("/edit/%d", postID), bytes.NewBufferString(editBody))
	editReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	editW := httptest.NewRecorder()
	r.ServeHTTP(editW, editReq)

	assert.Equal(t, http.StatusFound, editW.Code)
	fmt.Printf("Edit Post: %d OK\n", editW.Code)

	// Test 4: Add a comment to the post
	fmt.Println("\n=== Testing Add Comment ===")
	commentBody := "author=Commenter&email=commenter@example.com&content=This is a test comment."
	commentReq, _ := http.NewRequest("POST", fmt.Sprintf("/posts/%d/comments", postID), bytes.NewBufferString(commentBody))
	commentReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	commentW := httptest.NewRecorder()
	r.ServeHTTP(commentW, commentReq)

	assert.Equal(t, http.StatusFound, commentW.Code)
	fmt.Printf("Add Comment: %d OK\n", commentW.Code)

	// Test 5: Delete the post
	fmt.Println("\n=== Testing Delete Post ===")
	deleteBody := fmt.Sprintf("id=%d", postID)
	deleteReq, _ := http.NewRequest("POST", fmt.Sprintf("/delete/%d", postID), bytes.NewBufferString(deleteBody))
	deleteReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	deleteW := httptest.NewRecorder()
	r.ServeHTTP(deleteW, deleteReq)

	assert.Equal(t, http.StatusFound, deleteW.Code)
	fmt.Printf("Delete Post: %d OK\n", deleteW.Code)

	// Verify the post is deleted
	var deletedPost models.Post
	deleteErr := db.First(&deletedPost, postID).Error
	assert.Error(t, deleteErr)
	fmt.Println("Post deletion verified")
}

// TestMain sets up the test environment
func TestMain(m *testing.M) {
	// Set environment variables for database
	os.Setenv("DB_TYPE", "sqlite3")
	os.Setenv("DB_NAME", "test_blog.db")

	// Run all tests
	code := m.Run()

	// Clean up
	os.Remove("test_blog.db")

	// Exit with the test result code
	os.Exit(code)
}
