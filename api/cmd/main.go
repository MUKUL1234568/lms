package main

import (
	"library-management-api/config"
	"library-management-api/models"
	"library-management-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadConfig()

	// Connect to the database
	config.ConnectDB()
	db := config.DB

	// Run database migrations
	models.MigrateLibrary(db)
	models.MigrateUser(db)
	models.MigrateBook(db)
	models.MigrateRequestEvent(db)
	models.MigrateIssueRegistry(db)

	println("✅ Database Migrations Completed")

	// Initialize Gin Router
	router := gin.Default()

	// Register API Routes
	routes.AuthRoutes(router)
	routes.LibraryRoutes(router) // ✅ Register library routes
	routes.UserRoutes(router)    // ✅ User Routes
	routes.BookRoutes(router)
	routes.RequestRoutes(router)
	routes.IssueRegistryRoutes(router)

	// Start the server
	router.Run(":8080") // Run server on port 8080
}
