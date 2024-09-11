package main

import (
	"github.com/gin-gonic/gin"
	"go-auth-api/database"
	"go-auth-api/handlers"
)

func main() {
	router := gin.Default()

	// Connect to the database
	database.Connect()

	// Routes
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Start the server
	router.Run(":8080")
}
