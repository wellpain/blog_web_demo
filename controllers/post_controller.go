package controllers

import (
	"blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PostController handles post-related requests
type PostController struct {
	db *gorm.DB
}

// NewPostController creates a new PostController
func NewPostController(db *gorm.DB) *PostController {
	return &PostController{db: db}
}

// Index displays the welcome page
func (pc *PostController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": "我的博客",
	})
}

// AllPosts displays all blog posts
func (pc *PostController) AllPosts(c *gin.Context) {
	var posts []models.Post
	pc.db.Find(&posts)
	
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": "所有文章",
		"posts": posts,
		"showAllPosts": true,
	})
}

// Show displays a single post
func (pc *PostController) Show(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    post, err := models.GetPostByID(pc.db, uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    // 使用base.html作为基础模板，show.html作为内容块
    c.HTML(http.StatusOK, "base.html", gin.H{
        "Title": post.Title,
        "post": post,
        "showPost": true,
    })
}

// CreateForm displays the create post form
func (pc *PostController) CreateForm(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": "写文章",
		"showCreateForm": true,
	})
}

// Create creates a new post
func (pc *PostController) Create(c *gin.Context) {
	var post models.Post

	// Bind form data to post struct
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post in database
	if err := models.CreatePost(pc.db, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to index page
	c.Redirect(http.StatusFound, "/")
}

// EditForm displays the edit post form
func (pc *PostController) EditForm(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	post, err := models.GetPostByID(pc.db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title": "编辑文章",
		"post": post,
		"showEditForm": true,
	})
}

// Edit updates an existing post
func (pc *PostController) Edit(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Get existing post from database
	post, err := models.GetPostByID(pc.db, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Bind form data to post struct
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update post in database
	if err := models.UpdatePost(pc.db, &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to post page
	c.Redirect(http.StatusFound, "/posts/"+c.Param("id"))
}

// Delete deletes a post
func (pc *PostController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	// Delete post from database
	if err := models.DeletePost(pc.db, uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to index page
	c.Redirect(http.StatusFound, "/")
}

// AddComment adds a comment to a post
func (pc *PostController) AddComment(c *gin.Context) {
	postID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var comment models.Comment
	comment.PostID = uint(postID)

	// Bind form data to comment struct
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add comment to database
	if err := models.AddComment(pc.db, &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Redirect to post page
	c.Redirect(http.StatusFound, "/posts/"+c.Param("id"))
}