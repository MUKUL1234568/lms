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

		issueGroup.Use(middleware.AuthMiddleware())
		issueGroup.GET("/", middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.GetAllIssuedBooks)
		//issueGroup.GET("/user", middleware.RoleMiddleware("Reader"), controllers.GetIssuedBooksByReader)
	}
}
