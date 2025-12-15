package testdata

import (
	"blog/models"
	"blog/utils"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// Set default environment variables if .env file not found
		os.Setenv("DB_TYPE", "sqlite3")
		os.Setenv("DB_NAME", "blog.db")
	}

	// Initialize database connection
	utils.InitDB()

	// Create a test post
	post := models.Post{
		Title:   "测试文章",
		Content: "这是一篇测试文章的内容，用于展示博客应用的功能。",
		Author:  "测试作者",
	}

	// Save the test post to the database
	err = models.CreatePost(utils.DB, &post)
	if err != nil {
		panic(err)
	}

	println("测试文章已成功创建！")
}