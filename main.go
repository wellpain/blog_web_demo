package main

import (
	"blog/models"
	"blog/routes"
	"blog/utils"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env file not found, using system environment variables")
	}

	// Initialize database connection
	utils.InitDB()

	// Auto migrate database models
	utils.AutoMigrateDB(&models.Post{}, &models.Comment{})

	// Setup router
	r := routes.SetupRouter()

	// Get server port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	fmt.Printf("Server is running on port %s\n", port)
	fmt.Printf("Access the blog at http://localhost:%s\n", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}