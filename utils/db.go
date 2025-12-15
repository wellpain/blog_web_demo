package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	var err error

	// Get database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbName := os.Getenv("DB_NAME")

	// Connect to database based on type
	switch dbType {
	case "sqlite3":
		// Use modernc.org/sqlite which doesn't require CGO
		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
		// The connection string format remains the same
		DB, err = gorm.Open(sqlite.Open(dbName), gormConfig)
		if err != nil {
			log.Fatalf("Failed to connect to SQLite database: %v", err)
		}
		fmt.Println("Connected to SQLite database successfully")
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
	}
}

// AutoMigrateDB automatically migrates the database schema
func AutoMigrateDB(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}
	fmt.Println("Database migration completed successfully")
}
