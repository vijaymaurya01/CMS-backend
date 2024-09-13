package handlers

import (
	// "encoding/json"
	"go-auth-api/database"
	"go-auth-api/models"
	"go-auth-api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	UserName string `json:"userName"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	log.Println(input.Password)
	hashedPassword := utils.HashPassword(input.Password)
	log.Println("Entered password:", hashedPassword)

	// Save the user with the hashed password
	// Create the user with hashed password
	user := models.User{
		UserName: input.UserName,
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Address:  input.Address,
		Password: hashedPassword, // Store hashed password
		Role:     input.Role,
	}
	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}
	log.Println("Entered pass ", input.Password, " saved ", user.Password)
	// Compare the plain password with the hashed password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	// Generate JWT token
	token, err := utils.GenerateToken(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Generate JWT token if necessary...
	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}

type GetUserInput struct {
	Email string `json:"email" binding:"required"`
}

// GetUser retrieves user data based on the provided email
func GetUser(c *gin.Context) {
	var input GetUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Respond with user data
	c.JSON(http.StatusOK, gin.H{
		"userName": user.UserName,
		"name":     user.Name,
		"email":    user.Email,
		"phone":    user.Phone,
		"address":  user.Address,
		"role":     user.Role,
	})
}
func GetAllUsers(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with the list of users
	c.JSON(http.StatusOK, users)
}
