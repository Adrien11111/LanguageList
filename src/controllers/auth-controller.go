package controllers

import (
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"go-api/src/models"
	"go-api/src/repositories"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func Login(c *gin.Context) {
	var loginRequest models.RegisterRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	db := c.MustGet("DB").(*gorm.DB)
	userRepository := repositories.NewUserRepository(db)

	// Retrieve user by email
	user, err := userRepository.GetUserByEmail(loginRequest.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to retrieve user"})
		return
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})
	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return user without password and token
	response := gin.H{
		"id":    user.ID,
		"email": user.Email,
		"message": "Login successful",
		"token": token,
	}
	c.JSON(200, response)
}

func Register(c *gin.Context) {
	var registerRequest models.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	userRepository := repositories.NewUserRepository(db)


	// Check if the user already exists
	existingUser, err := userRepository.GetUserByEmail(registerRequest.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(500, gin.H{"error": "Failed to check existing user"})
		return
	}

	if existingUser != nil {
		c.JSON(409, gin.H{"error": "User already exists"})
		return
	}

	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Email:    registerRequest.Email,
		Password: string(passwordHash),
	}

	if err := userRepository.CreateUser(&user); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))


	// Return user without password
	response := gin.H{
		"id":    user.ID,
		"email": user.Email,
		"message": "User registered successfully",
		"token": token,
	}
	
	c.JSON(201, response)
}

func DeleteAccount(c *gin.Context) {
	user := c.MustGet("currentUser").(*models.User)
	db := c.MustGet("DB").(*gorm.DB)
	userRepository := repositories.NewUserRepository(db)

	if err := userRepository.DeleteUser(user.ID); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(204)
}