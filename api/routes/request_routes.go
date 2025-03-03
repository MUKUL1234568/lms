package routes

import (
	"library-management-api/controllers"
	"library-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// RequestRoutes registers book request routes
func RequestRoutes(router *gin.Engine) {
	requestGroup := router.Group("/requests")
	{
		// Protect all request routes with authentication
		requestGroup.Use(middleware.AuthMiddleware())

		// ✅ Only Readers can create book requests (Issue/Return)
		requestGroup.POST("/", middleware.RoleMiddleware("Reader"), controllers.CreateRequest)

		// ✅ Only LibraryAdmins can approve/reject requests
		requestGroup.PUT("/:id/approve", middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.ApproveRequest)

		// ✅ Readers can view their own requests
		requestGroup.GET("/user", middleware.RoleMiddleware("Reader"), controllers.GetUserRequests)

		// ✅ LibraryAdmins can view all requests
		requestGroup.GET("/allreq", middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.GetAllRequestsForAdmin)
	}
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzQwNDU4OTA2LCJyb2xlIjoiTGlicmFyeUFkbWluIiwidXNlcl9pZCI6NH0.7I2mgOp1aTKkYKrAYJVEsrNhZu_WUbqGHXBQ1AiWX1Q
