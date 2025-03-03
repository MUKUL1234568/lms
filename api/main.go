package main

import (
	"fmt"
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
	router.Use(CORSMiddleware())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173"}, // ✅ Allow frontend origin
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,           // Allow cookies and credentials to be sent
	// 	MaxAge:           12 * time.Hour, // Max Age for CORS preflight requests
	// }))

	routes.AuthRoutes(router)
	routes.LibraryRoutes(router) // ✅ Register library routes
	routes.UserRoutes(router)    // ✅ User Routes
	routes.BookRoutes(router)
	routes.RequestRoutes(router)
	routes.IssueRegistryRoutes(router)

	// Start the server
	router.Run(":8080") // Run server on port 8080
}

func CORSMiddleware() gin.HandlerFunc {
	fmt.Println("cors ")
	return func(c *gin.Context) {
		fmt.Println("header set")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}

}
