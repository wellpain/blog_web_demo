package models

import (
	"time"

	"gorm.io/gorm"
)

// Post represents a blog post
type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"size:255;not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Author    string         `gorm:"size:100;not null" json:"author"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	Comments  []Comment      `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}

// Comment represents a comment on a blog post
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PostID    uint           `json:"post_id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:100;not null" json:"email"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// GetAllPosts returns all posts
func GetAllPosts(db *gorm.DB) ([]Post, error) {
	var posts []Post
	err := db.Find(&posts).Error
	return posts, err
}

// GetPostByID returns a post by its ID
func GetPostByID(db *gorm.DB, id uint) (Post, error) {
	var post Post
	err := db.Preload("Comments").First(&post, id).Error
	// Ensure Comments is never nil, only empty slice
	if post.Comments == nil {
		post.Comments = []Comment{}
	}
	return post, err
}

// CreatePost creates a new post
func CreatePost(db *gorm.DB, post *Post) error {
	return db.Create(post).Error
}

// UpdatePost updates an existing post
func UpdatePost(db *gorm.DB, post *Post) error {
	return db.Save(post).Error
}

// DeletePost deletes a post by its ID
func DeletePost(db *gorm.DB, id uint) error {
	return db.Delete(&Post{}, id).Error
}

// AddComment adds a comment to a post
func AddComment(db *gorm.DB, comment *Comment) error {
	return db.Create(comment).Error
}