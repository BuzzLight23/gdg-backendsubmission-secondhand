package main

import (
	"gdg_backendsubmission_secondhand/config"
	"gdg_backendsubmission_secondhand/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to database
	config.ConnectDB()

	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	r.Run(":8080")
}
