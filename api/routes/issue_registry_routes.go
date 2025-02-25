package routes

import (
	"library-management-api/controllers"
	"library-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// IssueRegistryRoutes registers routes for issued books
func IssueRegistryRoutes(router *gin.Engine) {
	issueGroup := router.Group("/issueregistry")
	{
		// Protect all routes with authentication
		issueGroup.Use(middleware.AuthMiddleware())

		// ✅ LibraryAdmin can see all issued books
		issueGroup.GET("/admin", middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.GetAllIssuedBooks)

		// ✅ Readers can see their issued books
		issueGroup.GET("/user", middleware.RoleMiddleware("Reader"), controllers.GetIssuedBooksByReader)
	}
}
