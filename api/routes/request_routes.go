package routes

import (
	"library-management-api/controllers"
	"library-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// RequestRoutes registers book request routes
func RequestRoutes(router *gin.Engine) {
	requestGroup := router.Group("/request")
	{

		requestGroup.Use(middleware.AuthMiddleware())
		requestGroup.POST("/", middleware.RoleMiddleware("Reader"), controllers.CreateRequest)
		requestGroup.PUT("/:id", middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.ApproveRequest)
		//requestGroup.GET("/user", middleware.RoleMiddleware("Reader"), controllers.GetUserRequests)
		requestGroup.GET("/", middleware.RoleMiddlewareMultiple([]string{"LibraryAdmin", "Owner"}), controllers.GetAllRequestsForAdmin)
	}
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjoxNzQwNDU4OTA2LCJyb2xlIjoiTGlicmFyeUFkbWluIiwidXNlcl9pZCI6NH0.7I2mgOp1aTKkYKrAYJVEsrNhZu_WUbqGHXBQ1AiWX1Q
