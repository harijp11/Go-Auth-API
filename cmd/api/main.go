package main

import (
	// "database/sql"
	"log"
	"os"

	"Go-Auth-API/config"
	"Go-Auth-API/internal/initializer"

	"github.com/joho/godotenv"
	// "github.com/pressly/goose/v3"
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

	
	//2.5
	jwtConfig, err := config.LoadJWTConfig()
	if err != nil {
		log.Fatal("JWT configuration error:", err)
	}


	// 3. Close the database connection on exit
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}
	defer sqlDB.Close()

	// // 4. Run pending Goose migrations
	// if err := runMigrations(sqlDB); err != nil {
	// 	log.Fatal("Database migration failed:", err)
	// }

	// 5. Initialize the application and register routes
	router := initializer.InitializeApp(db,jwtConfig)

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

// func runMigrations(db *sql.DB) error {
// 	if err := goose.SetDialect("postgres"); err != nil {
// 		return err
// 	}

// 	return goose.Up(db, "migrations")
// }
