package main

import (
	"go-auth-api/database"
	"go-auth-api/handlers"
	"go-auth-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to the database
	database.Connect()

	// Routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/getuser", handlers.GetUser)
	router.GET("/users", middleware.AuthMiddleware("Admin"), handlers.GetAllUsers)

	// Start the server
	router.Run(":8080")
}
