package main

import (
	"gin-session-app/handlers"
	"gin-session-app/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Logging middleware for auth routes
	r.Use(middleware.LoggerMiddleware())

	r.POST("/login", handlers.Login)
	r.POST("/logout", handlers.Logout)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handlers.Profile)
	}

	r.Run(":8080")
}

