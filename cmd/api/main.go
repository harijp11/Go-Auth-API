
package main

import (
	"log"
	"os"

	"Go-Auth-API/config"
	"Go-Auth-API/internal/initializer"
	"Go-Auth-API/internal/models"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using system environment variables")
	}

	// 2. Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// 3. Close database connection when the application exits
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}
	defer sqlDB.Close()

	// 4. Run database migrations
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Database migration failed:", err)
	}

	// 5. Initialize the application and register routes
	router := initializer.InitializeApp(db)

	// 6. Get the server port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 7. Start the server
	log.Println("Server running on port:", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
